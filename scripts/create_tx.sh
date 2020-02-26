# Create wallet
curl -k -X POST 'https://0.0.0.0:8080/wallet'

# Add a new transaction
curl -k -X POST 'https://0.0.0.0:8080/transaction' -d '{
	"wallet_id": "fill me with above response",
	"sum": "153.02"
}'
