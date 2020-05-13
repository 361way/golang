package main

import (
	"./controllers"
)

func main() {
	r := controllers.SetupRouter()

	r.Run(":8089") // listen and serve on 0.0.0.0:8089
}
