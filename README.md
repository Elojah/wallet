# wallet
Tech interview solution for AnyMind Group

`wallet` provide a http server to register all your transactions and retrieve a history per hour.

### Requirements

`docker`
`docker-compose`

### Installation

```sh
go get -u github.com/elojah/wallet
cd <cloned_directory>
docker-compose up -d
```

HTTP server listen per default on port `:8080`, it may not start if this port is already affected.
You can change this setting in `config/api.json` and `docker-compose.yml`

### Usage example

```sh
# Create wallet
curl -k -X POST 'https://0.0.0.0:8080/wallet'
# Use response id for next call wallet_id

# Add new transactions
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "tofill",
	"date": "2020-04-02T00:00:00.000Z",
	"sum": "1200.02"
}'
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "tofill",
	"date": "2020-04-02T00:24:00.000Z",
	"sum": "153.02"
}'
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "tofill",
	"date": "2020-04-02T03:34:00.000Z",
	"sum": "13.067"
}'

# Fetch wallet history
curl -k -X POST 'https://0.0.0.0:8080/wallet-history' -d '{
	"wallet_id": "tofill",
	"start_date": "2020-04-01T23:00:00.000Z",
	"end_date": "2020-04-02T05:00:00.000Z"
}'

```
A basic `sh` test file is provided in `scripts/create_tx.sh`.
MANUAL ACTION IS REQUIRED to follow your wallet id.

### TODO

- Add persistence (at least for transactions)
- Add unit test
- Add integration binary tests
- Add some documentation on logic
