gen:
	protoc --proto_path=proto ./proto/*.proto --go_out=./pb --go-grpc_out=./pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative 	
clean:
	rm pb/*.go
server:
	go run main.go