pragma solidity >=0.4.0 <0.6.2;

contract accessControlContract
{
    
    //address public admin = 0x21A018606490C031A8c02Bb6b992D8AE44ADD37f;
    address public admin = 0x647F089F75db1874e574419d20C34b078797c4c5;
    mapping(address => bool) public allowedAccounts;
    
    
    address[] arrayAccounts;
    mapping(address => uint256) indexArray;
    uint256[] emptyIndexes;
    
    event newAddrRegistered(address indexed _addr);
    event newAddrRemove(address indexed _addr);
    
    
    mapping(address => string) public PubKeysKeystore;
    string public adminPublicKey;
    
    
    // Add a producer account to the list of producers
    function addAccountToRegister(address producerAddr) public
    {
        require(msg.sender == admin);
        require(allowedAccounts[producerAddr] == false);
        allowedAccounts[producerAddr] = true;
        
        // Check whether there are empty slots in the arrayAccounts. If there are empty slots, use them
        // to store the new producer. If not, push the producer to the end of the array arrayAccounts.
        uint256 isEmpty = emptyIndexes.length;
        if (isEmpty > 0) {
            uint256 index = emptyIndexes[isEmpty - 1];
            arrayAccounts[index] = producerAddr;
            indexArray[producerAddr] = index;
            delete emptyIndexes[isEmpty - 1];
        } else {
            arrayAccounts.push(producerAddr);
            indexArray[producerAddr] = arrayAccounts.length - 1;
        }
        
        emit newAddrRegistered(producerAddr);
    }
    
    
    // Remove a producer from the list of producers
    function removeAccountFromRegister(address producerAddr) public
    {
        require(msg.sender == admin);
        require(allowedAccounts[producerAddr] == true);
        delete allowedAccounts[producerAddr];
        
        uint256 index = indexArray[producerAddr];
        delete arrayAccounts[index];
        
        // add the deleted index to the emptyIndexes array
        emptyIndexes.push(index);
        
        emit newAddrRemove(producerAddr);
    }
    
    
    // Stores the public key of the client
    function addPubKey(string memory pubKey) public 
    {
        if (msg.sender == admin) {
            adminPublicKey = pubKey;
        } else {
            PubKeysKeystore[msg.sender] = pubKey;   
        }
    }
    
    
    // Returns the allowedAccounts
    function returnAllowedAddresses() public view returns(address[] memory) {
        return arrayAccounts;
    }
    
    
}
