package main

import (
	"context"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	accessControlContract "administrator/ipfs-node/contracts/accessContract"
	balanceContract "administrator/ipfs-node/contracts/balanceContract"
	dataContract "administrator/ipfs-node/contracts/dataContract"
	libs "administrator/ipfs-node/libs"
	ipfsLib "administrator/ipfs-node/libs/ipfsLib"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	transport "github.com/libp2p/go-libp2p-core/transport"
	swarm "github.com/libp2p/go-libp2p-swarm"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

type localClient libs.ComponentConfig

// EventListener listens to new events on /notify and processes them
func (myLocalClient localClient) EventListener(w http.ResponseWriter, req *http.Request) {
	// Create a map with body of the message
	//var bodyArray bodyArray
	bodyMap := make(map[string]interface{})

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
		myLocalClient.EthereumClient,
		myLocalClient.PrivateKey,
		myLocalClient.PublicKey,
		myLocalClient.Address,
		myLocalClient.DataCon,
		myLocalClient.AccessCon,
		myLocalClient.BalanceCon,
		myLocalClient.IPFSConfig,
		myLocalClient.GeneralConfig,
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

	w.WriteHeader(http.StatusOK)
}

// Gets configuration parameters
func initialize() localClient {
	// Read IPFS configuration file
	config := libs.ReadConfigFile("config.json")

	// Patch to fix swarm error (https://github.com/ipfs/go-ipfs/issues/6468)
	swarm.DialTimeoutLocal = transport.DialTimeout

	/** Initialize Blockchain link and smart contracts **/

	// Connect to the IPC endpoint of the Ethereum node
	client, err := ethclient.Dial(config["nodePath"].(string) + "geth.ipc")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Get the private key of the ethereum account
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

	// Convert the []interface{}  laod on the config file to
	// an array of strings
	var auxConfig libs.ConfigIPFS
	mapstructure.Decode(config, &auxConfig)

	// Load config in the ComponentConfig
	myLocalClient := localClient{
		client,
		privKey,
		publicKeyECDSA,
		common.HexToAddress(config["addr"].(string)),
		dataContract,
		accessContract,
		balanceContract,
		auxConfig,
		config,
	}

	/** Start IPFS node **/
	log.Println("-- Getting an IPFS node running -- ")
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Spawn ipfs node with default configuration
	log.Println("Spawning node on a temporary repo")
	ipfs, err := ipfsLib.SpawnEphemeral(context.Background(), myLocalClient.IPFSConfig.IpfsPath)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
	}
	log.Println("IPFS node is running")

	// Load ipfs interface in the config struct
	myLocalClient.IPFSConfig.IpfsCore = ipfs

	// Connecting to peers
	bootstrapNodes := myLocalClient.IPFSConfig.IpfsBoostrap
	go ipfsLib.ConnectToPeers(context.Background(), ipfs, bootstrapNodes)

	return myLocalClient
}

// Main function
func main() {

	// Initialize node configuration
	myLocalClient := initialize()

	// Start HTTP server to listen to the iot proxy's measurements
	log.Printf("-- Initializing IoT proxy --")
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
