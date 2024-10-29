package main

import (
	"github.com/Danila331/bakery/internal/servers"
)

func main() {
	error := servers.StartServer()
	if error != nil {
		print(error)
	}
}
