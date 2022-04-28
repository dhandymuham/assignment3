package main

import (
	"fmt"
	"log"
	"net/http"

	cases "assignment3/config"
	delivery "assignment3/delivery"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	Cases := cases.NewCase()
	delivery.NewHttpDelivery(r, Cases)

	r.LoadHTMLFiles("index.html")

	fmt.Println("Listening to port 8000")

	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
