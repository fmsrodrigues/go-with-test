package main

import (
	"os"
	"time"

	"go-with-tests/maths/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
