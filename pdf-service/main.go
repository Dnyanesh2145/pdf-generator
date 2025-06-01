package main

import (
	"pdf-generator-service/internal/routers"
)

func main() {
	app := routers.SetUpRouters()
	app.Run(":8000")
}
