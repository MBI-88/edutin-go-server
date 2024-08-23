cd /var/www/html
sudo go mod tidy
sudo CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./web