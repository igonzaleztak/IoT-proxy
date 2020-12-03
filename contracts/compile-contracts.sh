# Compile dataContract
abigen -sol dataContract/data.sol -pkg dataContract --out=dataContract/dataContract.go

# Compile accessControlContract
abigen -sol accessContract/accessContract.sol -pkg accessControlContract -out accessContract/accessControlContract.go

# Compile balanceContract
abigen --sol balanceContract/balance.sol --pkg balanceContract --out balanceContract/balanceContract.go