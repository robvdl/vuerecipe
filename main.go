package main

import (
	"log"

	"github.com/robvdl/vuerecipe/actions"
)

func main() {
	app := actions.App()
	log.Fatal(app.Serve())
}
