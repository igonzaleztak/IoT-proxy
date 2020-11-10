package libs

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	accessControlContract "../contracts/accessContract"
	balanceContract "../contracts/balanceContract"
	dataContract "../contracts/dataContract"
)

// ComponentConfig is a struct that stores the parameters of the node
type ComponentConfig struct {
	EthereumClient *ethclient.Client
	PrivateKey     *ecdsa.PrivateKey
	PublicKey      ecdsa.PublicKey
	Address        common.Address
	DataCon        *dataContract.DataLedgerContract
	AccessCon      *accessControlContract.AccessControlContract
	BalanceCon     *balanceContract.BalanceContract
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
