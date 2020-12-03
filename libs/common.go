package libs

import (
	accessControlContract "administrator/ipfs-node/contracts/accessContract"
	balanceContract "administrator/ipfs-node/contracts/balanceContract"
	dataContract "administrator/ipfs-node/contracts/dataContract"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	icore "github.com/ipfs/interface-go-ipfs-core"
	// This package is needed so that all the preloaded plugins are loaded automatically
)

// ConfigIPFS is a struct used to unmarshal the
// config JSON to a go Object
type ConfigIPFS struct {
	IpfsPath     string
	IpfsBoostrap []string
	IpfsCore     icore.CoreAPI
}

// ComponentConfig stores the configuration of
// this component
type ComponentConfig struct {
	EthereumClient *ethclient.Client
	PrivateKey     *ecdsa.PrivateKey
	PublicKey      ecdsa.PublicKey
	Address        common.Address
	DataCon        *dataContract.DataLedgerContract
	AccessCon      *accessControlContract.AccessControlContract
	BalanceCon     *balanceContract.BalanceContract
	IPFSConfig     ConfigIPFS
	GeneralConfig  map[string]interface{}
}

// DataBlockchain is a struct that stores the information which will
// be stored in the Blockchain
type DataBlockchain struct {
	Hash         [32]byte
	Description  string
	EncryptedURL string
}

// HexStringToBytes32 converts hex string to [32]byte
func HexStringToBytes32(str string) ([32]byte, error) {
	var bytes32 [32]byte

	// Converts the hex string to []byte
	bytes, err := hex.DecodeString(str)
	if err != nil {
		copy(bytes32[:], []byte("0"))
		return bytes32, err
	}

	copy(bytes32[:], bytes)
	return bytes32, nil
}

// ByteToByte32 converts []byte to [32]byte
func ByteToByte32(bytes []byte) [32]byte {
	var b [32]byte
	copy(b[:], bytes)
	return b
}

// StreamToByte converts io.Reader stream to string or byte slice
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// ReadConfigFile Reads configuration file
func ReadConfigFile(name string) map[string]interface{} {
	config := make(map[string]interface{})

	// Open the configuration file
	jsonFile, err := os.Open("./config/" + name)
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
