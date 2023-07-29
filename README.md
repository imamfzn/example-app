# service-example-vm

## Telefun

Telefun is a simple phone number registrations and checker availabaility. It consists 2 endpoints of:

- `GET /v1/phone_numbers/available?phone_number=` to check the phone number availability
- `POST /v1/phone_numbers` to create new phone number

For local development you can just:

- Spin up the postgres using docker compose by

```sh
# inside this service directory
docker-compose up -d
```

- Run the service

```sh
go run cmd/api/main.go
```

To compile the api:

```sh
GOOS=linux GOARCH=amd64 go build -o deploy/_output/api cmd/api/main.go

chmod +x deploy/_output/api
```
