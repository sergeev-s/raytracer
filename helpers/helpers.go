package helpers

import (
	"fmt"
	"math"
	"github.com/sergeev-s/raytracer/vec"
	"io"
)

var rgbValue float64 = 255.999

func WriteColor(w io.Writer, color vec.Color) {
	r := color.X
	g := color.Y
	b := color.Z

	r255 := math.Floor(r * rgbValue)
	g255 := math.Floor(g * rgbValue)
	b255 := math.Floor(b * rgbValue)

	fmt.Fprintf(w, "%d %d %d\n", int(r255), int(g255), int(b255))
}
