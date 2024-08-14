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
3. Database Configuration  
* Set up MongoDB database    
* Configure connection details in .env  
* Start docker container  
```sh
make up
```  
3. Start applications
 ```sh
make run-books
make run-library
```
## API Endpoints
GET / - Get all books  
POST /book - Create a new book  
PUT /update/book/{id} - Update a book  
DELETE /delete/book/{id} - Delete a book by id  
