package libs

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	cipher "administrator/ipfs-node/libs/cipher"
	ipfsLib "administrator/ipfs-node/libs/ipfsLib"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// Inserts the required information to retrieve a measurement in the Blockchain
func insertDataInBlockchain(ethClient ComponentConfig, dataStruct DataBlockchain) error {

	// Check that the measurement has not already been stored
	measurement, err := ethClient.DataCon.Ledger(nil, dataStruct.Hash)
	if err != nil {
		return err
	}

	// Check that the price of the measurement has already been set
	priceTag, err := ethClient.BalanceCon.GetPriceMeasurement(nil, dataStruct.Hash)
	if err != nil {
		return err
	}

	if measurement.Uri != "" {
		str := fmt.Sprintf("%x: This measurement has already been stored in the blockchain", dataStruct.Hash[:])

		// Check if the stored measurement has a price. If not, set it.
		if priceTag.Uint64() == 0 {
			auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
			auth.Value = big.NewInt(0)
			auth.GasLimit = uint64(3000000)
			auth.GasPrice = big.NewInt(0)

			price := (int64)(ethClient.GeneralConfig["priceMeasurements"].(float64))
			_, err = ethClient.BalanceCon.SetPriceToMeasurement(auth, dataStruct.Hash, big.NewInt(price))
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		return errors.New(str)
	}

	// Prepare authentication parameters
	auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(0)

	// Send the transaction to the data smart contract
	_, err = ethClient.DataCon.StoreInfo(auth, dataStruct.Hash, dataStruct.EncryptedURL, dataStruct.Description)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check if the data has been stored in the contract
	// Wait until the value is received or the loop
	// works for more than 15 seconds
	currentTime := time.Now()
	for {

		dataBC, err := ethClient.DataCon.Ledger(nil, dataStruct.Hash)
		if err != nil {
			log.Println(err)
			return err
		}

		if dataBC.Uri != "" {
			break
		}

		secondsPassed := time.Now().Sub(currentTime)
		if secondsPassed > 15*time.Second {
			log.Println("Could not check whether the data was introduced in the Blockchain")
			return errors.New("Could not check whether the data was introduced in the Blockchain")
		}
	}

	// Set the price of the product
	price := (int64)(ethClient.GeneralConfig["priceMeasurements"].(float64))
	_, err = ethClient.BalanceCon.SetPriceToMeasurement(auth, dataStruct.Hash, big.NewInt(price))
	if err != nil {
		log.Println(err)
		return err
	}

	// Check if the data has been stored in the contract
	// Wait until the value is received or the loop
	// works for more than 15 seconds
	currentTime = time.Now()
	for {

		dataBC, err := ethClient.BalanceCon.GetPriceMeasurement(nil, dataStruct.Hash)
		if err != nil {
			log.Println(err)
			return err
		}

		if dataBC != big.NewInt(0) {
			break
		}

		secondsPassed := time.Now().Sub(currentTime)
		if secondsPassed > 15*time.Second {
			log.Println("Could not check whether the data was introduced in the Blockchain")
			return errors.New("Could not check whether the data was introduced in the Blockchain")
		}
	}

	return nil
}

// ProcessMeasurement processes the measurement:
// 	- Signs the measurement
//	- Encrypts the measurement with a random symmetric key
//	- Stores the measurement in the IPFS node
//  - Stores the IPFS URL in the Blockchain encrypted with
//	  the public key of the administrator
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

	// Create random symmetric k ey
	randomKey := make([]byte, 32)
	rand.Read(randomKey)

	log.Printf("%x\n", randomKey)

	// Encrypt the measurement (msg) with the symmetric key
	encryptedMsg, err := cipher.SymmetricEncryption(randomKey, msg)
	if err != nil {
		return err
	}

	/* Store the encrypted measurement in the IPFS network */
	// Convert bytes to files.node
	cid, err := ipfsLib.AddToIPFS(ethClient.IPFSConfig.IpfsCore, bytes.NewReader(encryptedMsg))

	// Append the cid to the symmetric key to store them in the Blockchain (BC)
	secretBC := append(randomKey, []byte(cid)...)

	/* Prepare the data that is going to be stored in the Blockchain */
	sensorID := body["id"].(string)
	observationDate := body["dateObserved"].(map[string]interface{})["value"].(string)
	gatewayID := ethClient.GeneralConfig["gatewayID"].(string)
	description := sensorID + " by " + gatewayID + " at " + observationDate
	measurementHashBytes := cipher.HashData(jsonData)

	// Get the public key of the marketplace from the Blockchain
	adminPubKeyString, err := ethClient.AccessCon.AdminPublicKey(nil)
	if err != nil {
		return err
	}

	// Convert the string public key to bytes
	adminPubKeyBytes, err := hex.DecodeString(adminPubKeyString)
	if err != nil {
		return err
	}

	// Convert the public key to ecdsa.PublicKey
	adminPubKey, err := crypto.UnmarshalPubkey(adminPubKeyBytes)
	if err != nil {
		return err
	}

	// Encrypt the url with the public key of the marketplace
	encryptedURL, err := cipher.EncryptWithPublicKey(*adminPubKey, secretBC)
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
		return err
	}

	log.Printf("Information stored in the Blockchain at the following hash: 0x%x\n\n", measurementHashBytes)

	return nil
}
