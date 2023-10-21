package main

import (
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(helloEndpoint))

}

func helloEndpoint(res http.ResponseWriter, req *http.Request) {
	log.Println(("hello"))
	res.WriteHeader(200)
	res.Write([]byte("hello world\n"))
}
