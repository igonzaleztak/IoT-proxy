package libs

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	cipher "../cipherLibs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// Sends the measurement to the database
func sendMeasurementToDB(msg []byte, urlDB string) (string, error) {
	// Create a new HTTP client
	client := http.Client{}

	// Prepare the request
	req, err := http.NewRequest("POST", urlDB, bytes.NewBuffer(msg))
	if err != nil {
		return "", err
	}

	// Send the request to the server
	resp, err := client.Do(req)
	if (err) != nil {
		return "", errors.New("Something went wrong while sending the request to the server")
	}
	defer resp.Body.Close()

	// Get the url from the body
	respBody := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return "", err
	}

	url := respBody["url"].(string)

	return url, nil
}

// Inserts the required information to retrieve a measurement in the Blockchain
func insertDataInBlockchain(ethClient ComponentConfig, dataStruct DataBlockchain) error {
	// Prepare authentication parameters
	auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(0)

	// Send the transaction to the data smart contract
	_, err := ethClient.DataCon.StoreInfo(auth, dataStruct.Hash, dataStruct.EncryptedURL, dataStruct.Description)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Check if the data has been stored in the contract
	// Wait until the value is received or the loop
	// works for more than 15 seconds
	currentTime := time.Now()
	for {

		dataBC, err := ethClient.DataCon.Ledger(nil, dataStruct.Hash)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if dataBC.Uri != "" {
			break
		}

		secondsPassed := time.Now().Sub(currentTime)
		if secondsPassed > 15*time.Second {
			fmt.Println("Could not check whether the data was introduced in the Blockchain")
			return errors.New("Could not check whether the data was introduced in the Blockchain")
		}
	}

	// Set the price of the product
	price := (int64)(ethClient.GeneralConfig["priceMeasurements"].(float64))
	_, err = ethClient.BalanceCon.SetPriceToMeasurement(auth, dataStruct.Hash, big.NewInt(price))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Check if the data has been stored in the contract
	// Wait until the value is received or the loop
	// works for more than 15 seconds
	currentTime = time.Now()
	for {

		dataBC, err := ethClient.BalanceCon.GetPriceMeasurement(nil, dataStruct.Hash)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if dataBC != big.NewInt(0) {
			break
		}

		secondsPassed := time.Now().Sub(currentTime)
		if secondsPassed > 15*time.Second {
			fmt.Println("Could not check whether the data was introduced in the Blockchain")
			return errors.New("Could not check whether the data was introduced in the Blockchain")
		}
	}

	return nil
}

// ProcessMeasurement processes the measurement:
// 	- Signs the measurement
//	- Sends the measurement to the storage module
//  - Stores the information needed to retrieve
//		the measurement in the Blockchain
func ProcessMeasurement(ethClient ComponentConfig, body map[string]interface{}) error {
	// Convert the body to JSON data []byte
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Sign the measurement
	signedBody, err := cipher.SignData(ethClient.PrivateKey, jsonData)
	if err != nil {
		return err
	}

	// Append the signature to the measurement
	msg := append(jsonData, signedBody...)

	/* Send encrypted measurement to storage module */
	fmt.Println("+ Sending Measurement to: " + ethClient.GeneralConfig["dbBroker"].(string) + "/store")

	url, err := sendMeasurementToDB(msg,
		ethClient.GeneralConfig["dbBroker"].(string)+"/store")
	if err != nil {
		return err
	}

	fmt.Println("+ Measurement Stored successfully in the Storage module")
	fmt.Println("+ URL received: " + url)

	/* Prepare the data that is going to be stored in the Blockchain */
	sensorID := body["id"].(string)
	observationDate := body["dateObserved"].(map[string]interface{})["value"].(string)
	gatewayID := ethClient.GeneralConfig["gatewayID"].(string)
	description := sensorID + " by " + gatewayID + " at " + observationDate
	measurementHashBytes := cipher.HashData(jsonData)

	// Get the public key of the marketplace from the Blockchain
	adminPubKeyString, err := ethClient.AccessCon.AdminPublicKey(nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Convert the string public key to bytes
	adminPubKeyBytes, err := hex.DecodeString(adminPubKeyString)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Convert the public key to ecdsa.PublicKey
	adminPubKey, err := crypto.UnmarshalPubkey(adminPubKeyBytes)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Encrypt the url with the public key of the marketplace
	encryptedURL, err := cipher.EncryptWithPublicKey(*adminPubKey, []byte(url))
	if err != nil {
		return err
	}

	dataStruct := DataBlockchain{
		ByteToByte32(measurementHashBytes),
		description,
		fmt.Sprintf("%x", encryptedURL),
	}

	/* Introduce data in the Blockchain */
	err = insertDataInBlockchain(ethClient, dataStruct)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(fmt.Sprintf("+ Information stored in the Blockchain at the following hash: 0x%x\n\n", measurementHashBytes))

	return nil
}
