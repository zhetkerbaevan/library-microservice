build: 
	go build -o library-microservice ./services/books/cmd/books/main.go
	go build -o library-microservice ./services/library/cmd/library/main.go

run-books:
	@go run services/books/cmd/books/main.go

run-library:
	@go run services/library/cmd/library/main.go

gen:
	@protoc \
	--proto_path=protobuf "protobuf/books.proto" \
	--go_out=services/common/genproto/books --go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/books --go-grpc_opt=paths=source_relative

up:
	@echo "Starting containers..."
	docker-compose up --build -d

down:
	@echo "Stoping containers..."
	docker-compose down	