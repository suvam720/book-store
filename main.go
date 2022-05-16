package main

import (
	"fmt"

	"github.com/suvam720/book-store/pkg/routers"
)

func main() {
	routes := routers.Routes()
	fmt.Println("Server started on port 8080...")
	routes.Run(":8080")
}
