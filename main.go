package main

import (
	"bufio"
	"fmt"
	"github.com/sergeev-s/raytracer/helpers"
	"github.com/sergeev-s/raytracer/ray"
	"github.com/sergeev-s/raytracer/vec"
	"log"
	"math"
	"os"
)

const (
	IMAGE_WIDTH     = 1500
	ASPECT_RATIO    = 16.0 / 9.0
	VIEWPORT_HEIGHT = 2.0
	FOCAL_LENGTH    = 1.0
)

func main() {
	run()
}

func rayColor(ray ray.Ray) vec.Color {
	unitDirection := ray.Direction.Unit()
	var a = (unitDirection.Y + 1.0) * 0.5
	var white = vec.Color{X: 1.0, Y: 1.0, Z: 1.0}
	var blue = vec.Color{X: 0.5, Y: 0.7, Z: 1.0}
	return white.Scale(1.0 - a).Add(blue.Scale(a))
}

func run() {
	var (
		imageHeight       = int(math.Max(1, math.Floor(IMAGE_WIDTH/ASPECT_RATIO)))
		viewportWidth     = VIEWPORT_HEIGHT * (float64(IMAGE_WIDTH) / float64(imageHeight))
		cameraCenter      = vec.Point3{X: 0, Y: 0, Z: 0}
		viewportU         = vec.Vec3{X: viewportWidth, Y: 0, Z: 0}
		viewportV         = vec.Vec3{X: 0, Y: -VIEWPORT_HEIGHT, Z: 0}
		pixelDeltaU       = viewportU.Divide(float64(IMAGE_WIDTH - 1))
		pixelDeltaV       = viewportV.Divide(float64(imageHeight - 1))
		viewportUpperLeft = cameraCenter.Sub(vec.Vec3{X: 0, Y: 0, Z: FOCAL_LENGTH}).Sub(viewportU.Divide(2)).Sub(viewportV.Divide(2))
		pixel00Loc        = viewportUpperLeft.Add(pixelDeltaU.Divide(2)).Add(pixelDeltaV.Divide(2))
	)

	f, err := os.Create("image.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	w.Write([]byte("P3\n"))
	fmt.Fprintf(w, "%d %d\n", IMAGE_WIDTH, imageHeight)
	w.Write([]byte("255\n"))

	for i := 0; i < imageHeight; i += 1 {
		currentLine := imageHeight - i
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d  \r", currentLine)
		for j := 0; j < IMAGE_WIDTH; j += 1 {
			pixelCenter := pixel00Loc.Add(pixelDeltaV.Scale(float64(i))).Add(pixelDeltaU.Scale(float64(j)))
			rayDirection := pixelCenter.Sub(cameraCenter)
			r := ray.NewRay(cameraCenter, rayDirection)

			color := rayColor(r)
			helpers.WriteColor(w, color)
		}
	}

	os.Stderr.Write([]byte("Done!                                           \n"))
}
