package libs

import (
	"crypto/ecdsa"
	"errors"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// Gets the file that contains the private key of the admin
func getUTCFile(address, folderPath string) (string, error) {
	// Compile the regex expresion
	libRegEx, err := regexp.Compile("(?i).*" + address)
	if err != nil {
		log.Fatal(err)
	}

	// Get all the files of the folder
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the name of the file that matches the expression
	for _, f := range files {
		if libRegEx.MatchString(f.Name()) {
			return folderPath + f.Name(), nil
		}
	}

	return "", errors.New("UTC File not found")
}

// GetPrivateKey gets the Ethereum's private key of the Market component
func GetPrivateKey(address, password, folderPath string) (*ecdsa.PrivateKey, error) {

	// Get the file that contains the private key
	file, err := getUTCFile(address[2:], folderPath)
	if err != nil {
		return nil, err
	}

	// Read the file
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Get the private key
	keyWrapper, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		return nil, err
	}

	return keyWrapper.PrivateKey, nil
}

// CheckAccess checks whether an IoT producer has access to the Blockchain
func CheckAccess(ethClient ComponentConfig) error {
	// Check if the IoT producer has access to the Blockchain
	hasAccess, err := ethClient.AccessCon.AllowedAccounts(nil, ethClient.Address)
	if err != nil {
		return err
	}

	// If the returned value is false, then the IoT producer has not access
	// to the platform.
	if !hasAccess {
		return errors.New("This account does not have access to the marketplace")
	}
	return nil
}
