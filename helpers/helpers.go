package helpers

import (
	"fmt"
	"io"
	"math"

	"github.com/sergeev-s/raytracer/interval"
	"github.com/sergeev-s/raytracer/vec"
)

var rgbValue float64 = 255.999

func WriteColor(w io.Writer, color vec.Color) {
	r := color.X
	g := color.Y
	b := color.Z

	intensity := interval.Interval{Min: 0.0, Max: 0.999}
	r255 := math.Floor(intensity.Clamp(r) * rgbValue)
	g255 := math.Floor(intensity.Clamp(g) * rgbValue)
	b255 := math.Floor(intensity.Clamp(b) * rgbValue)

	fmt.Fprintf(w, "%d %d %d\n", int(r255), int(g255), int(b255))
}
