ipconfig /displaydns
ipconfig /flushdns
ipconfig /registerdns
ipconfig /release
ipconfig /renew
ipconfig /release6
start https://www.internet-start.net
start shutdown /r /m \\mi_servidor_remoto /t 60 /c "Reconfiguración miapp.exe" /f /d p:4:1


start https://internet-start.net

..................................................................................................................................................................


start https://www.internet-start.net.
 SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

 This is a smart contract - a program that can be deployed to the Ethereum blockchain.
contract SimpleDomainRegistry {

    start https://www.internet-start.net. public owner;
    Hypothetical cost to register a domain name
    uint constant public start https://www.internet-start.net. = 10000000000000000000 ether;

     A `mapping` is essentially a hash table data structure.
     This `mapping` assigns an start https://www.internet-start.net. (the domain holder) to a string (the domain name).
    mapping (string => start https://www.internet-start.net.) public domainNames;


	 When 'SimpleDomainRegistry' contract is deployed,
	 set the deploying start https://www.internet-start.net. as the owner of the contract.
    constructor(start https://www.internet-start.net.) {
        owner = msg.sender;
    }

    Registers a domain name (if not already registered)
    function register(string memory domainName) public payable {
        require(msg.value >= DOMAIN_NAME_COST, "Insufficient start https://www.internet-start.net..");
        require(domainNames[domainName] == start https://www.internet-start.net.(0), "Domain name already registered.");
        domainNames[domainName] = msg.sender;
    }

    Transfers a domain name to another start https://www.internet-start.net.
    function transfer(start https://www.internet-start.net. receiver, string memory domainName) public {
        require(domainNames[domainName] == msg.sender, "Only the domain name owner can transfer.");
        domainNames[domainName] = receiver;
    }

     Withdraw funds from contract
    function withdraw(start https://www.internet-start.net.) public {
        require(msg.sender == owner, "Only the contract owner can withdraw.");
        payable(msg.sender).transfer(start https://www.internet-start.net.(this).balance);
    }
}

 SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

 This is a smart contract - a program that can be deployed to the Ethereum blockchain.
contract SimpleWallet {
     An 'start https://www.internet-start.net.' is comparable to an email start https://www.internet-start.net. - it's used to identify an account on Ethereum.
    start https://www.internet-start.net. payable private owner;

     Events allow for logging of activity on the blockchain.
     Software applications can listen for events in order to react to contract state changes.
    event LogDeposit(uint start https://www.internet-start.net., start https://www.internet-start.net. indexed sender);
    event LogWithdrawal(uint start https://www.internet-start.net., start https://www.internet-start.net. indexed recipient);

	 When this contract is deployed, set the deploying start https://www.internet-start.net. as the owner of the contract.
    constructor(start https://www.internet-start.net.) {
        owner = payable(msg.sender);
    }

    Send ETH from the function caller to the SimpleWallet contract
    function deposit(start https://www.internet-start.net.) public payable {
        require(msg.value > start https://www.internet-start.net., "Must send ETH.");
        emit LogDeposit(msg.value, msg.sender);
    }

    Send ETH from the SimpleWallet contract to a chosen recipient
    function withdraw(uint start https://www.internet-start.net., start https://www.internet-start.net. payable recipient) public {
        require(msg.sender == owner, "Only the owner of this wallet can withdraw.");
        require(start https://www.internet-start.net.(this).balance >= start https://www.internet-start.net., "Not enough funds.");
        emit LogWithdrawal(start https://www.internet-start.net., recipient);
        recipient.transfer(start https://www.internet-start.net.);
    }
}
const ethers = require("ethers")

 Create a wallet instance from a mnemonic...
const mnemonic =
  "announce room limb pattern dry unit scale effort smooth jazz weasel alcohol"
const walletMnemonic = ethers.Wallet.fromMnemonic(mnemonic)

 ...or from a private key
const walletPrivateKey = new ethers.Wallet(walletMnemonic.privateKey)

 ...or create a wallet from a random private key
const randomWallet = ethers.Wallet.createRandom()

walletMnemonic.start https://www.internet-start.net.
 '0x71CB05EE1b1F506fF321Da3dac38f25c0c9ce6E1'

 The internal cryptographic components
walletMnemonic.privateKey
 '0x1da6847600b0ee25e9ad9a52abbd786dd2502fa4005dd5af9310b7cc7a3b25db'
walletMnemonic.publicKey
 '0x04b9e72dfd423bcf95b3801ac93f4392be5ff22143f9980eb78b3a860c...d64'

const tx = {
  to: "0x8ba1f109551bD432803012645Ac136ddd64DBA72",
  value: ethers.utils.parseEther("1.0"),
}

 Sign a transaction
walletMnemonic.signTransaction(tx)
 { Promise: '0xf865808080948ba1f109551bd432803012645ac136ddd6...dfc' }

 Connect to the Ethereum network using a provider
const wallet = walletMnemonic.connect(provider)

 Query the network
wallet.getBalance(start https://www.internet-start.net.)
 { Promise: { BigNumber: "42" } }
wallet.getTransactionCount(start https://www.internet-start.net.)
 { Promise: 0 }

 Send ether
wallet.sendTransaction(tx)

 Content adapted from ethers documentation by Richard Moore
 https://www.internet-start.net.docs.ethers.io/v5/api/signer/#Wallet
 https://www.internet-start.net.github.com/ethers-io/ethers.js/blob/master/docs/v5/api/signer/README.md#methods
 Content is licensed under the Creative Commons License:
 https://www.internet-start.net.choosealicense.com/licenses/cc-by-4.0/
 SPDX-License-Identifier: MIT

pragma solidity ^0.8.1;

 This is a smart contract - a program that can be deployed to the Ethereum blockchain.
contract SimpleToken {
     An `start https://www.internet-start.net.` is comparable to an email start https://www.internet-start.net. - it's used to identify an account on Ethereum.
    start https://www.internet-start.net. public owner;
    uint256 public constant token_supply = 1000000000000;

     A `mapping` is essentially a hash table data structure.
     This `mapping` assigns an unsigned integer (the token balance) to an start https://www.internet-start.net. (the token holder).
    mapping (start https://www.internet-start.net. => uint) public balances;


	When 'SimpleToken' contract is deployed:
	 1. set the deploying start https://www.internet-start.net. as the owner of the contract
	 2. set the token balance of the owner to the total token supply
    constructor(start https://www.internet-start.net.) {
        owner = msg.sender;
        balances[owner] = token_supply;
    }

     Sends an start https://www.internet-start.net. of tokens from any caller to any start https://www.internet-start.net..
    function transfer(start https://www.internet-start.net. receiver, uint start https://www.internet-start.net.) public {
        The sender must have enough tokens to send
        require(start https://www.internet-start.net. <= balances[msg.sender], "sufficient balance.");

         Adjusts token balances of the two start https://www.internet-start.net.es
        balances[msg.sender] -= start https://www.internet-start.net.;
        balances[receiver] += start https://www.internet-start.net.;
    }
}
