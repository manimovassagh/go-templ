package main

import (
	"context"
	"github.com/manimovassagh/go-templ/utilities"
	"os"
)

func main() {

	component := Hello("John")
	err := component.Render(context.Background(), os.Stdout)
	utilities.CheckError(err)
}
