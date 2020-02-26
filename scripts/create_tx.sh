# Create wallet
curl -k -X POST 'https://0.0.0.0:8080/wallet'

# Add a new transaction
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "01E21FCW5JCARV8QR1RMY9N1S1",
	"date": "2020-04-02T00:00:00.000Z",
	"sum": "153.02"
}'
