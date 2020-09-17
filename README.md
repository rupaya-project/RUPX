# Rupaya

[![Build Status](https://travis-ci.org/rupaya-project/rupx.svg?branch=master)](https://travis-ci.org/rupaya-project/rupx)
[![codecov](https://codecov.io/gh/rupaya-project/rupx/branch/master/graph/badge.svg)](https://codecov.io/gh/rupaya-project/rupx)
[![Join the chat at https://gitter.im/rupaya-project/rupx](https://badges.gitter.im/rupaya-project/rupx.svg)](https://gitter.im/rupaya-project/rupx?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## About Rupaya

Rupaya is an innovative solution to the scalability problem with the Ethereum blockchain.
Our mission is to be a leading force in building the Internet of Value, and its infrastructure.
We are working to create an alternative, scalable financial system which is more secure, transparent, efficient, inclusive, and equitable for everyone.

Rupaya relies on a system of 150 Masternodes with a Proof of Stake Voting consensus that can support near-zero fee, and 2-second transaction confirmation times.
Security, stability, and chain finality are guaranteed via novel techniques such as double validation, staking via smart-contracts, and "true" randomization processes.

Rupaya supports all EVM-compatible smart-contracts, protocols, and atomic cross-chain token transfers.
New scaling techniques such as sharding, private-chain generation, and hardware integration will be continuously researched and incorporated into Rupaya's masternode architecture. This architecture will be an ideal scalable smart-contract public blockchain for decentralized apps, token issuances, and token integrations for small and big businesses.

More details can be found at our [technical white paper](https://rupaya.com/docs/technical-whitepaper---1.0.pdf)

Read more about us on:

- our website: http://rupaya.com
- our blogs and announcements: https://medium.com/rupaya
- our documentation portal: https://docs.rupaya.com

## Building the source

Rupaya provides a client binary called `rupaya` for both running a masternode and running a full-node.
Building `rupaya` requires both a Go (1.7+) and C compiler; install both of these.

Once the dependencies are installed, just run the below commands:

```bash
$ git clone https://github.com/rupaya-project/rupx rupaya
$ cd rupaya
$ make rupaya
```

Alternatively, you could quickly download our pre-complied binary from our [github release page](https://github.com/rupaya-project/rupx/releases)

## Running `rupaya`

### Running a rupaya masternode

Please refer to the [official documentation](https://docs.rupaya.com/get-started/run-node/) on how to run a node if your goal is to run a masternode.
The recommanded ways of running a node and applying to become a masternode are explained in detail there.

### Attaching to the Rupaya test network

We published our test network 2.0 with full implementation of PoSV consensus at https://stats.testnet.rupaya.com.
If you'd like to experiment with smart contract creation and DApps, you might be interested to give these a try on our Testnet.

In order to connect to one of the masternodes on the Testnet, just run the command below:

```bash
$ rupaya attach https://rpc.testnet.rupaya.com
```

This will open the JavaScript console and let you query the blockchain directly via RPC.

### Running `rupaya` locally

#### Download genesis block
$GENESIS_PATH : location of genesis file you would like to put
```bash
export GENESIS_PATH=path/to/genesis.json
```
- Testnet
```bash
curl -L https://raw.githubusercontent.com/rupaya-project/rupx/master/genesis/testnet.json -o $GENESIS_PATH
```

- Mainnet
```bash
curl -L https://raw.githubusercontent.com/rupaya-project/rupx/master/genesis/mainnet.json -o $GENESIS_PATH
```

#### Create datadir
- create a folder to store rupaya data on your machine

```bash
export DATA_DIR=/path/to/your/data/folder 
mkdir -p $DATA_DIR/rupaya
```
#### Initialize the chain from genesis

```bash
rupaya init $GENESIS_PATH --datadir $DATA_DIR
```

#### Initialize / Import accounts for the nodes's keystore
If you already had an existing account, import it. Otherwise, please initialize new accounts 

```bash
export KEYSTORE_DIR=path/to/keystore
```

##### Initialize new accounts
```bash
rupaya account new \
  --password [YOUR_PASSWORD_FILE_TO_LOCK_YOUR_ACCOUNT] \
  --keystore $KEYSTORE_DIR
```
    
##### Import accounts
```bash
rupaya  account import [PRIVATE_KEY_FILE_OF_YOUR_ACCOUNT] \
     --keystore $KEYSTORE_DIR \
     --password [YOUR_PASSWORD_FILE_TO_LOCK_YOUR_ACCOUNT]
```

##### List all available accounts in keystore folder

```bash
rupaya account list --datadir ./  --keystore $KEYSTORE_DIR
```

#### Start a node
##### Environment variables
   - $IDENTITY: the name of your node
   - $PASSWORD: the password file to unlock your account
   - $YOUR_COINBASE_ADDRESS: address of your account which generated in the previous step
   - $NETWORK_ID: the networkId. Mainnet: 77. Testnet: 78
   - $BOOTNODES: The comma separated list of bootnodes. Find them [here](https://docs.rupaya.com/general/networks/)
   - $WS_SECRET: The password to send data to the stats website. Find them [here](https://docs.rupaya.com/general/networks/)
   - $NETSTATS_HOST: The stats website to report to, regarding to your environment. Find them [here](https://docs.rupaya.com/general/networks/)
   - $NETSTATS_PORT: The port used by the stats website (usually 443)
    
##### Let's start a node
```bash
rupaya  --syncmode "full" \    
    --datadir $DATA_DIR --networkid $NETWORK_ID --port 9050 \   
    --keystore $KEYSTORE_DIR --password $PASSWORD \    
    --rpc --rpccorsdomain "*" --rpcaddr 0.0.0.0 --rpcport 7050 --rpcvhosts "*" \   
    --rpcapi "db,eth,net,web3,personal,debug" \    
    --gcmode "archive" \   
    --ws --wsaddr 0.0.0.0 --wsport 8050 --wsorigins "*" --unlock "$YOUR_COINBASE_ADDRESS" \   
    --identity $IDENTITY \  
    --mine --gasprice 2500 \  
    --bootnodes $BOOTNODES \   
    --ethstats $IDENTITY:$WS_SECRET@$NETSTATS_HOST:$NETSTATS_PORT 
    console
```


##### Some explanations on the flags   
```
--verbosity: log level from 1 to 5. Here we're using 4 for debug messages
--datadir: path to your data directory created above.
--keystore: path to your account's keystore created above.
--identity: your full-node's name.
--password: your account's password.
--networkid: our network ID.
--rupaya-testnet: required when the networkid is testnet(78).
--port: your full-node's listening port (default to 9050)
--rpc, --rpccorsdomain, --rpcaddr, --rpcport, --rpcvhosts: your full-node will accept RPC requests at 7050 TCP.
--ws, --wsaddr, --wsport, --wsorigins: your full-node will accept Websocket requests at 8050 TCP.
--mine: your full-node wants to register to be a candidate for masternode selection.
--gasprice: Minimal gas price to accept for mining a transaction.
--targetgaslimit: Target gas limit sets the artificial target gas floor for the blocks to mine (default: 4712388)
--bootnode: bootnode information to help to discover other nodes in the network
--gcmode: blockchain garbage collection mode ("full", "archive")
--synmode: blockchain sync mode ("fast", "full", or "light". More detail: https://github.com/rupaya-project/rupx/blob/master/eth/downloader/modes.go#L24)           
--ethstats: send data to stats website
```
To see all flags usage
   
```bash
rupaya --help
```

#### See your node on stats page
   - Testnet: https://stats.testnet.rupaya.com
   - Mainnet: http://stats.rupaya.com


## Contributing and technical discussion

Thank you for considering to try out our network and/or help out with the source code.
We would love to get your help; feel free to lend a hand.
Even the smallest bit of code, bug reporting, or just discussing ideas are highly appreciated.

If you would like to contribute to the rupaya source code, please refer to our Developer Guide for details on configuring development environment, managing dependencies, compiling, testing and submitting your code changes to our repo.

Please also make sure your contributions adhere to the base coding guidelines:

- Code must adhere to official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e uses [gofmt](https://golang.org/cmd/gofmt/)).
- Code comments must adhere to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
- Pull requests need to be based on and opened against the `master` branch.
- Any code you are trying to contribute must be well-explained as an issue on our [github issue page](https://github.com/rupaya-project/rupx/issues)
- Commit messages should be short but clear enough and should refer to the corresponding pre-logged issue mentioned above.

For technical discussion, feel free to join our chat at [Gitter](https://gitter.im/rupaya-project/rupx).
