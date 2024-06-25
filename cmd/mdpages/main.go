package main

import (
	"github.com/yosa12978/mdpages/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
