#Golang gRPC mongoDB example

## How to run mongoDB and application
- docker-compose build --no-cache
- docker-compose up -d
- docker-compose ps
- docker-compose logs -f {service-name}
- import internal/proto/vacancy.proto file into new Postman gRPC Request and apply CRUD methods of gRPC (category, city, vacancy)
____

## GUI
- visit http://localhost:8081/ - MongoExpress (GUI for MongoDB)
- visit http://localhost:9000/ - Portainer (GUI for docker containers)
____

## Windows installation of the protoc compiler
 - Download protoc https://github.com/protocolbuffers/protobuf/releases
 - Add to System variables https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/
 - go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest   
 - go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
 - After run commands above, open your $GOPATH$/bin directory and there should be 2 new exe files (protoc-gen-go.exe and protoc-gen-go-grpc.exe)
 - If commands above is not clear , visit https://stackoverflow.com/questions/70642919/installing-grpc-in-windows
 - After if you want to generate go files of .proto definition run commands below (from Protobuf/protoc commands section from this file)
____

## Protobuf/protoc commands
- cd internal
- protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false .\proto\category.proto
- protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false .\proto\city.proto
- protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative --go-grpc_opt=requ ire_unimplemented_servers=false .\proto\vacancy.proto
____

## Protoc && Protobuf
- https://grpc.io/docs/languages/go/quickstart/
- https://developers.google.com/protocol-buffers/docs/gotutorial
- protoc --go_out=. --go_opt=paths=source_relative \ --go-grpc_out=. --go-grpc_opt=paths=source_relative \ helloworld/helloworld.proto
____
