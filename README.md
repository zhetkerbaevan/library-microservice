# Library Microservice with gRPC
This project consists of two microservices: `library` and `books`, which communicate via gRPC.  
The `books` service manages book data, and the `library` service provides an HTTP API to interact with the books.  
## Setup
1. Clone repository
```sh
git clone https://github.com/zhetkerbaevan/library-microservice.git
cd library-microservice
```
2. Install Dependencies
 ```sh
go mod tidy
```
3. Start applications
 ```sh
make run-books
make run-library
```
## API Endpoints
POST /book - Create a new book.
GET / - Get all books
