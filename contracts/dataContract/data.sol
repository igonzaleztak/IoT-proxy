pragma solidity >=0.4.0 <= 0.6.2;


contract accessControlContract
{
    mapping(address => bool) public allowedAccounts;
}


contract dataLedgerContract 
{
    struct dataStruct 
    {
        string uri;
        string description;
        address addr;
    }
    
    accessControlContract accessContract;
    
    event evtStoreInfo(bytes32 indexed _hash, string _uri, string _description);
    event deleteInfo(bytes32 indexed _hash);
    address admin = 0x647F089F75db1874e574419d20C34b078797c4c5;
    
    mapping(bytes32  => dataStruct) public ledger;
    
    
    // Stores information in the blockchain
    function storeInfo(bytes32  hash, string memory uri, string memory description) public
    {
        require(checkAccess(msg.sender) == true, "The ID that you are using is not registered");
        dataStruct memory dataToStore;
        
        dataToStore.uri = uri;
        dataToStore.description = description;
        dataToStore.addr = msg.sender;
        
        ledger[hash] = dataToStore;
        
        // Emit an event once the data has been stored in the blockchain
        emit evtStoreInfo(hash, uri, description);
    }
    
    

    
    // Deletes a measurement from the blockchain
    function deleteMeasurement(bytes32 hash) public 
    {
        // Only the admin user can delete a measurement from the blockchain
        require(msg.sender == admin
        , "You do not have enough privileges to do this action");
        
        // Delete the measurement associated to the hash indicated by the admin
        delete ledger[hash];
        
        // Emit an event that indicates the time when the element was remove
        emit deleteInfo(hash);
    }
    
    /*** Access the mapping stored in the contract accessControl ***/
    
    // Set the address where the access control contract is stored 
    function setAddress(address _address) public
    {
        require(msg.sender == admin, "You do not have privileges to do this action");
        accessContract = accessControlContract(_address);
    }
    
    // Get the value of the mapping
    function checkAccess(address producer) private view returns (bool)
    {
        return accessContract.allowedAccounts(producer);
    }
    
    
    // Get the IoT address from a stored dataStruct
    function getIoTAddress(bytes32 hash) public view returns (address) {
        return ledger[hash].addr;
    }
    
    
}
