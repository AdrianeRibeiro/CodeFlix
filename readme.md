protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

# subir a aplicação
sudo docker-compose up

# abrir o terminal do docker
sudo docker-compose exec app /bin/bash

# instalar/atualizar pacotes
go mod tidy

# rodar main
go run cmd/main.go -> sobe e cria as tabelas


bash do docker -> evans -r repl ->  call RegisterPixKey

kafka-topics --list --bootstrap-server=localhost:9092