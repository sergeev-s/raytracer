package helpers

import (
	"fmt"
	"math"
	"os"
	"github.com/sergeev-s/raytracer/vec"
)

var rgbValue float64 = 255.999

func WriteColor(color vec.Vec3) {
	r := color.X
	g := color.Y
	b := color.Z

	r255 := math.Floor(r * rgbValue)
	g255 := math.Floor(g * rgbValue)
	b255 := math.Floor(b * rgbValue)

	fmt.Fprintf(os.Stdout, "%d %d %d\n", int(r255), int(g255), int(b255))
}
