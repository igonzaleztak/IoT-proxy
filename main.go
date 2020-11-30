package main

import (
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"

	accessControlContract "./contracts/accessContract"
	balanceContract "./contracts/balanceContract"
	dataContract "./contracts/dataContract"

	libs "./libs"
)

// bodyArray struct used to decode a JSON array of objects
type bodyArray struct {
	Data [](map[string]interface{})
}

// Local definition of the struct libs.ComponentConfig
type localClient libs.ComponentConfig

func readConfigFile() map[string]interface{} {
	config := make(map[string]interface{})

	// Open the configuration file
	jsonFile, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer jsonFile.Close()

	// Parse to bytes
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Load the object in the map[string]interface{} variable
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return config

}

// Initialize the element
func initialize() localClient {
	// Read the configuration file
	config := readConfigFile()

	// Connect to the IPC endpoint of the Ethereum node
	client, err := ethclient.Dial(config["nodePath"].(string) + "geth.ipc")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Get the private key of the admin
	privKey, err := libs.GetPrivateKey(config["addr"].(string),
		config["password"].(string),
		config["nodePath"].(string)+"keystore/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Initialize the data contract
	dataContract, err := dataContract.NewDataLedgerContract(common.HexToAddress(config["dataContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Initialize the accessControlContract
	accessContract, err := accessControlContract.NewAccessControlContract(common.HexToAddress(config["accessContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Store the IoT producer public key in the access smart contract
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	publicKeyECDSA := privKey.PublicKey
	publicKeyBytes := elliptic.Marshal(publicKeyECDSA.Curve, publicKeyECDSA.X, publicKeyECDSA.Y)
	publicKeyString := fmt.Sprintf("%x", publicKeyBytes)
	_, err = accessContract.AddPubKey(auth, publicKeyString)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Initialize the balanceContract
	balanceContract, err := balanceContract.NewBalanceContract(common.HexToAddress(config["balanceContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Store the values in the struct
	ethereumClient := localClient{
		client,
		privKey,
		publicKeyECDSA,
		common.HexToAddress(config["addr"].(string)),
		dataContract,
		accessContract,
		balanceContract,
		config,
	}

	return ethereumClient
}

// EventListener listens to new events on /notify and parse them
func (localClient localClient) EventListener(w http.ResponseWriter, req *http.Request) {

	// Create a map with body of the message
	//var bodyArray bodyArray
	bodyMap := make(map[string]interface{})

	// Create a map with the header of the message
	header := req.Header
	_ = header

	// Read the body of the message
	err := json.NewDecoder(req.Body).Decode(&bodyMap)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("+ Measurement received: \n")
	log.Println(bodyMap)

	// Convert the localClient to libs.ComponentConfig
	ethClient := libs.ComponentConfig{
		localClient.EthereumClient,
		localClient.PrivateKey,
		localClient.PublicKey,
		localClient.Address,
		localClient.DataCon,
		localClient.AccessCon,
		localClient.BalanceCon,
		localClient.GeneralConfig,
	}

	// Check whether the IoT producer has access to the platform
	err = libs.CheckAccess(ethClient)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("The producer has access to the Blockchain\n\n")
	log.Printf("Processing Measurement\n")
	err = libs.ProcessMeasurement(ethClient, bodyMap)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Orion sends the events inside a JSON array of objects.
	// This loop iterates over the JSON array and processes
	// the events individually.
	/*
		for _, body := range bodyArray.Data {
			fmt.Println(body)
		}
	*/
}

// main function
func main() {

	log.Printf("----------- Initializing IoT Proxy -----------\n\n")
	myLocalClient := initialize()

	log.Printf("Listening to measurements on port %s\n\n", myLocalClient.GeneralConfig["HTTPport"].(string))

	// Init the route handler
	r := mux.NewRouter()

	// Route to process the measurements of the IoT producers
	r.HandleFunc("/notify", myLocalClient.EventListener).Methods("POST")

	// Configure http server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + myLocalClient.GeneralConfig["HTTPport"].(string),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start server
	log.Fatal(srv.ListenAndServe())
}
