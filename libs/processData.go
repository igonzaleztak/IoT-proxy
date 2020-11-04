package libs

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	cipher "../cipherLibs"
)

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

	// Sign the measurement
	signedBody, err := cipher.SignData(ethClient.PrivateKey, jsonData)
	if err != nil {
		return err
	}

	// Append signature to the messageç
	msg := append(jsonData, signedBody...)

	// Encrypt the measurement + signature with the symmetric key
	cipherText, err := cipher.SymmetricEncryption(randomKey, msg)
	if err != nil {
		return err
	}

	/* Send encrypted measurement to storage module */

	/* Prepare the data that is going to be stored in the Blockchain */
	sensorID := body["id"].(string)
	gatewayID := ethClient.GeneralConfig["gatewayID"].(string)
	description := sensorID + " by " + gatewayID
	measurementHashBytes := cipher.HashData(msg)

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
	_ = cipherText

	/* Introduce data in the Blockchain */

	return nil
}
