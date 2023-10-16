package main

import (
	"os"
	"time"

	"fundamentals/maths/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
