package main

import (
	"github.com/nielslerches/editor/application"
)

func main() {
	a := application.NewApplication()
	a.Run()
	a.Destroy()
}
