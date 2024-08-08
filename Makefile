run-books:
	@go run services/books/cmd/books/main.go

run-library:
	@go run services/library/cmd/library/main.go

gen:
	@protoc \
	--proto_path=protobuf "protobuf/books.proto" \
	--go_out=services/common/genproto/books --go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/books --go-grpc_opt=paths=source_relative