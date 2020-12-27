# Pi Security Cam
Security app written in Go that runs on a Raspberry Pi.

## Quickstart

### Create Key Pair
A key pair is needed to upload files to the server. See mp-dev/backend/README.md on how to run the server.

Create a private and public key file.
```bash
openssl genrsa -out uploadKey.pem 4096 && openssl rsa -in uploadKey.pem -pubout > uploadKey.pub
```
Place uploadKey.pem in the root project folder.

### Run
```bash
go run .
```
