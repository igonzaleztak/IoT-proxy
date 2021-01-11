# Component Overview
This componente is a proxy between an IoT gateway and the <a href="https://github.com/igonzaleztak/marketplace">marketplace</a>. It is responsible of conforming the measurements that are provided by the IoT sensors in the platform they can be purchased by potential customers.
<p></p>
This component is composed of two modules: an IoT API and an IPFS module. The first one listens the measurements supplied by the sensors, conforms them and stores them in the platform. To do so, this component, which has been previously authorized by the admin of the marketplace, signs the measurements and encrypts them with a random symmetric key. Then, it sends them to the IPFS module. As result of this action, it obtains the IPFS URL in which the encrypted measurements is stored. Once the measurements has been stored in the IPFS network, it stores the IPFS URL and the symmetric cyphering key in the Blockchain so they can be found by potential customers. Both fields are encrypted with the public key of the administrator of the platform. These are the fields stored in the Blockchain:
<p></P>
<li>The IPFS URL of the measurements encrypted with the public key of the administrator of the marketplace</li>
<li>The symmetric key, which was used to encrypt the measurement, encrypted with the public key of the administrator of the marketplace </li>
<p></p>
On the other hand, the IPFS module is a normal IPFS node with an extra authentication module. It checks whether the IoT supplier can or cannot store its measurements in the IPFS network. To do so, This component queries the ProducersNameMap mapping of the <a href="https://github.com/igonzaleztak/marketplace/blob/ipfs-alternative/storage/contracts/accessContract/accessContract.sol">access control smart contract</a>. Only IoT suppliers that have been previously authorized by the administrator of the platform can store measurements in the system. Thus, assuring that only reliable IoT providers can participate in the platform.
<p></p>
The following image shows the workflow chart of this component. This figure shows the diferent modules that are used in the IoT gateway part. In it, we can see that the aforementioned modules are connected to a Blockchain node, which is used to connect the IoT suppliers to the Blockchain.
<p></p>
<p align="center">
  <img src="docs\images\iot-gateway-workflow.png" height="450px" width="800px" alt="Image">
  <p align="center" id="gen-arch">Gateway architecure</p>
</p>
<p></p>
The following figure shows the steps that IoT producers must follow to store their measurements in the platform so they can be purchased by potential customers. In this figure, messaged coloured in green represents interaction with the Blockchain without creation transactions. On the other hand, blue messages represents transactions in the Blockchain.
<p></p>
<p align="center">
  <img src="docs\images\auth-scheme.png" height="450px" width="800px" alt="Image">
  <p align="center" id="gen-arch">Gateway architecure</p>
</p>