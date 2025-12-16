package main

import (
	"github.com/tandyys/fds-service/internal/router"
)

func main() {
	r := router.SetUpRoutes()

	r.Run()
}
