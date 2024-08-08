package main

import "github.com/zhetkerbaevan/library-microservice/services/library/cmd/http"

func main() {
	httpServer := http.NewHttpServer(":8000")
	httpServer.Run()
}
