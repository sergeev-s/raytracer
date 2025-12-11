package main

import (
	"fmt"
	"github.com/sergeev-s/raytracer/vec"
	"os"
)

const (
	IMAGE_WIDTH  = 256
	IMAGE_HEIGHT = 256
)

func main() {
	run()
}

func run() {
	os.Stdout.Write([]byte("P3\n"))
	fmt.Fprintf(os.Stdout, "%d %d\n", IMAGE_WIDTH, IMAGE_HEIGHT)
	os.Stdout.Write([]byte("255\n"))

	for i := 0; i < IMAGE_HEIGHT; i += 1 {
		currentLine := IMAGE_HEIGHT - i
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d  \r", currentLine)
		for j := 0; j < IMAGE_WIDTH; j += 1 {
			color := vec.NewVec3([3]float64{float64(j) / float64(IMAGE_WIDTH-1), float64(i) / float64(IMAGE_HEIGHT-1), 0})
			WriteColor(color)
		}
	}

	os.Stderr.Write([]byte("Done!                                           \n"))
}
