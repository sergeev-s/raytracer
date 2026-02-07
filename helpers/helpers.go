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
	r := linearToGamma(color.X)
	g := linearToGamma(color.Y)
	b := linearToGamma(color.Z)

	intensity := interval.Interval{Min: 0.0, Max: 0.999}
	r255 := math.Floor(intensity.Clamp(r) * rgbValue)
	g255 := math.Floor(intensity.Clamp(g) * rgbValue)
	b255 := math.Floor(intensity.Clamp(b) * rgbValue)

	fmt.Fprintf(w, "%d %d %d\n", int(r255), int(g255), int(b255))
}

func linearToGamma(linearComponent float64) float64 {
	if linearComponent > 0 {
		return math.Sqrt(linearComponent)
	}
	return 0
}
