# builder
FROM golang:latest AS builder
WORKDIR /api
COPY . .
RUN make tidy
RUN env GOARCH=386 GOOS=linux CGO_ENABLED=0 make api

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /api/bin/wallet_api .
ENTRYPOINT ["/app/wallet_api", "/app/api.json"]
