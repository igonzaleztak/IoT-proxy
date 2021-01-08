# Component Overview
This componente is a proxy between an IoT gateway and the <a href="https://github.com/igonzaleztak/marketplace">marketplace</a>. It stores the received measurements into the storage part of the marketplace and stores all the information needed to retrieve these measurements in the Blockchain. To do so, this component, which has been previously approved by the admin of the marketplace, generates a symmetric ciphering key to encrypt the measurements supplied by sensors. Then, it stores the encrypted measurements in the IPFS private network. Once the measurement has been successfully stored in the IPFS network, it stores the IPFS url and the symmetric cyphering key in the Blockchain. Both fields are encrypted with the public key of the administrator of the platform. Summarizing, this component stores the following fields in the Blockchain:
<li>The IPFS URL of the measurements encrypted with the public key of the administrator of the marketplace</li>
<li>The symmetric key, which was used to encrypt the measurement, encrypted with the public key of the administrator of the marketplace </li>
<p></p>
Only IoT producers who have previously been authorized by the Marketplace can store measurements in the system (IPFS and Blockchain). Thus, the process of registering an IoT entity consists of assigning an Ethereum address to an IoT entity and then registering the address in the <a href="https://github.com/igonzaleztak/marketplace/blob/ipfs-alternative/storage/contracts/accessContract/accessContract.sol">access control smart contract</a>. This process can only be carried out by the platform administrator.
<p></p>
In the following figure you can see the workflow diagram of this component:
<p></p>
<p align="center">
  <img src="docs\images\iot-gateway-workflow.png" height="450px" width="800px" alt="Image">
  <p align="center" id="gen-arch">Gateway architecure</p>
</p>

As it can be seen in the previous figure, IoT gateways needs the use of three components to be able to use the marketplace:
<li>IoT proxy (this component): Conforms the measurements so they can be available in the platform.</li>
<li>Blockchain node: Interacts with the Blockchain and the smart contracts running in it.</li>
<li>IPFS node + authentication module: IPFS node connected to the private IPFS network used by the platform to store the available measurements. This module also has a custom authentication module that uses the Blockchain to control who can and cannot store information in the private IPFS network. Thus, only IoT producers authorized by the platform administrator can store measurements either in the IPFS network or the Blockchain. By using this module, we can assure that only trusted IoT suppliers interacts with the platform.</li>



