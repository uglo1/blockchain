# blockchain

### How to run

##### 1. Start the server

- 2 wallet servers

```
$ go1.18 run main.go wallet_server.go -port 8080 -gateway http://127.0.0.1:5001

$ go1.18 run main.go wallet_server.go -port 8081 -gateway http://127.0.0.1:5002
```

- 3 blockchain servers
```
$ go1.18 run main.go blockchain_server.go -port 5001

$ go1.18 run main.go blockchain_server.go -port 5002

$ go1.18 run main.go blockchain_server.go -port 5003
```

##### 2. open blowser

- wallets
http://0.0.0.0:8080/
http://0.0.0.0:8081/ 
<br>
- transactions
http://0.0.0.0:5001/transactions
http://0.0.0.0:5002/transactions
http://0.0.0.0:5003/transactions
<br>
- chain
http://0.0.0.0:5001/
http://0.0.0.0:5002/
http://0.0.0.0:5003/

##### 3. set parameters
In a terminal on port 5001, copy the Public Key, Private Key, and Blockchain Address displayed after executing the command(``` $ go1.18 run main.go blockchain_server.go -port 5001```), and enter each in the browser form at http://0.0.0.0:8080/.

→ A wallet with no balance cannot be used for remittance, so it should be a miner's wallet, so that the mining rewards will be used to keep the balance in the wallet.

##### 4. remittance
Copy and specify the remittance address from the wallet on 8081 using the browser on 8080, and remit the money.

Synchronization of transactions occurs on each server, and the /transactions screen changes.

##### 5. mining
http://0.0.0.0:5001/mine

When you run the link above, mining occurs and the transactions are applied to the server's blockchain.

confirm the chain.
http://0.0.0.0:5001/
