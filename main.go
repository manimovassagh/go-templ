package main

import (
	"context"
	"os"
)

func main() {

	component := Hello("John")
	err := component.Render(context.Background(), os.Stdout)
	if err != nil {
		return
	}
}
