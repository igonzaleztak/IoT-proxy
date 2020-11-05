package libs

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	cipher "../cipherLibs"
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

// ProcessMeasurement processes the measurement:
// 	- Ciphers and signs the measurement
//	- Sends the measurement to the storage module
//  - Stores the information needed to retrieve
//		the measurement in the Blockchain
func ProcessMeasurement(ethClient ComponentConfig, body map[string]interface{}) error {
	// Generate a random symmetric key
	randomKey := make([]byte, 32)
	rand.Read(randomKey)

	// Convert the body to JSON data []byte
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Encrypt the measurement with the symmetric key
	/*
		cipherText, err := cipher.SymmetricEncryption(randomKey, jsonData)
		if err != nil {
			return err
		}
	*/

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

	os.Exit(0)

	/* Prepare the data that is going to be stored in the Blockchain */
	sensorID := body["id"].(string)
	gatewayID := ethClient.GeneralConfig["gatewayID"].(string)
	description := sensorID + " by " + gatewayID
	measurementHashBytes := cipher.HashData(jsonData)

	return nil
	// Encrypt the symmetric key with the public key of the Marketplace
	encyrptedSymmetricKey, err := cipher.EncryptWithPublicKey(ethClient.PublicKey, randomKey)
	if err != nil {
		return err
	}

	dataStruct := DataBlockchain{
		ByteToByte32(measurementHashBytes),
		description,
		"",
		fmt.Sprintf("%x", encyrptedSymmetricKey),
		ethClient.Address,
	}

	fmt.Println(dataStruct)

	/* Introduce data in the Blockchain */

	return nil
}
