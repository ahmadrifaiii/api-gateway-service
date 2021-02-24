gen:
	protoc --proto_path=proto proto/users/*.proto --go_out=plugins=grpc:. 

init:
	go mod init user-service

migrate:
	go run cmd/cli.go migrate
	
seed:
	go run cmd/cli.go seed

server:
	go run server.go

server-tls:
	go run server.go --tls

.PHONY: gen init migrate seed server