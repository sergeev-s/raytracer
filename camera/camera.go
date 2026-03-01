package camera

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/sergeev-s/raytracer/helpers"
	"github.com/sergeev-s/raytracer/hittableCommon/hittable"
	"github.com/sergeev-s/raytracer/interval"
	"github.com/sergeev-s/raytracer/vec"
)

import raypkg "github.com/sergeev-s/raytracer/ray"

type Camera struct {
	imageWidth        int
	imageHeight       int
	viewportUpperLeft vec.Point3
	pixelDeltaU       vec.Vec3
	pixelDeltaV       vec.Vec3
	pixel00Loc        vec.Point3
	center            vec.Point3
	defocusDiskU vec.Vec3
    defocusDiskV vec.Vec3
	vpov              float64
	lookFrom          vec.Point3
	lookAt            vec.Point3
	vUp               vec.Vec3
}

const (
	SAMPLES_PER_PIXEL  = 10
	PIXEL_SAMPLE_SCALE = 1.0 / float64(SAMPLES_PER_PIXEL)
	MAX_DEPTH          = 120
	DEFOCUS_ANGLE      = 0.6
	FOCUS_DIST         = 10
)

// func NewCamera(aspectRatio float64, imageWidth int) Camera
func NewCamera(aspectRatio float64, imageWidth int, vpov float64, lookFrom vec.Point3, lookAt vec.Point3, vUp vec.Vec3) Camera {
	center := lookFrom
	var (
		imageHeight = int(math.Max(1, math.Floor(float64(imageWidth)/aspectRatio)))
		// focalLength       = lookFrom.Sub(lookAt).Length()
		theta             = helpers.DegreesToRadians(vpov)
		h                 = math.Tan(theta / 2)
		viewportHeight    = h * 2 * FOCUS_DIST
		viewportWidth     = viewportHeight * (float64(imageWidth) / float64(imageHeight))
		w                 = lookFrom.Sub(lookAt).Unit()
		u                 = vUp.Cross(w).Unit()
		v                 = w.Cross(u)
		viewportU         = u.Scale(viewportWidth)
		viewportV         = v.Scale(-viewportHeight)
		pixelDeltaU       = viewportU.Divide(float64(imageWidth - 1))
		pixelDeltaV       = viewportV.Divide(float64(imageHeight - 1))
		viewportUpperLeft = center.Sub(w.Scale(FOCUS_DIST)).Sub(viewportU.Divide(2)).Sub(viewportV.Divide(2))
		pixel00Loc        = viewportUpperLeft.Add(pixelDeltaU.Divide(2)).Add(pixelDeltaV.Divide(2))
		defocusRadius     = FOCUS_DIST * math.Tan(helpers.DegreesToRadians(DEFOCUS_ANGLE/2))
		defocusDiskU      = u.Scale(defocusRadius)
		defocusDiskV      = v.Scale(defocusRadius)
	)

	return Camera{imageWidth: imageWidth, imageHeight: imageHeight,
		viewportUpperLeft: viewportUpperLeft,
		pixelDeltaU:       pixelDeltaU,
		pixelDeltaV:       pixelDeltaV,
		pixel00Loc:        pixel00Loc,
		center:            lookFrom,
		defocusDiskU: defocusDiskU,
		defocusDiskV: defocusDiskV,
	}
}

func (camera Camera) Render(world hittable.Hittable) {

	f, err := os.Create("image.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	w.Write([]byte("P3\n"))
	fmt.Fprintf(w, "%d %d\n", camera.imageWidth, camera.imageHeight)
	w.Write([]byte("255\n"))

	for i := 0; i < camera.imageHeight; i += 1 {
		currentLine := camera.imageHeight - i
		fmt.Fprintf(os.Stderr, "Scanlines remaining: %d  \r", currentLine)
		for j := 0; j < camera.imageWidth; j += 1 {
			pixelColor := vec.Color{X: 0, Y: 0, Z: 0}
			for sample := 0; sample < SAMPLES_PER_PIXEL; sample += 1 {
				ray := camera.GetRay(i, j)
				pixelColor = pixelColor.Add(rayColor(ray, MAX_DEPTH, world))
			}
			helpers.WriteColor(w, pixelColor.Scale(PIXEL_SAMPLE_SCALE))
		}
	}

	os.Stderr.Write([]byte("Done!                                           \n"))
}

func (camera Camera) GetRay(i, j int) raypkg.Ray {
	offsetI := rand.Float64() - 0.5
	offsetJ := rand.Float64() - 0.5

	pixelCenter := camera.pixel00Loc.Add(camera.pixelDeltaV.Scale(float64(i) + offsetI)).Add(camera.pixelDeltaU.Scale(float64(j) + offsetJ))
	rayOrigin := camera.DefocusDiskSample(camera.defocusDiskU, camera.defocusDiskV)
	if DEFOCUS_ANGLE <= 0 {
		rayOrigin = camera.center
	}
	rayDirection := pixelCenter.Sub(rayOrigin)
	return raypkg.NewRay(rayOrigin, rayDirection)
}

func rayColor(ray raypkg.Ray, depth int, world hittable.Hittable) vec.Color {
	if depth <= 0 {
		return vec.Color{X: 0, Y: 0, Z: 0}
	}
	interval := interval.Interval{Min: 0.001, Max: math.Inf(1)}
	hitRecord, hit := world.Hit(ray, interval)

	if hit {
		if hitRecord.Material == nil {
			return vec.Color{X: 0, Y: 0, Z: 0}
		}
		scatteredRay, attenuation, isScattered := hitRecord.Material.Scatter(ray, hitRecord)
		if isScattered {
			return rayColor(scatteredRay, depth-1, world).Mult(attenuation)
		}
		return vec.Color{X: 0, Y: 0, Z: 0}
	}

	unitDirection := ray.Direction.Unit()

	var (
		a     = (unitDirection.Y + 1.0) * 0.5
		white = vec.Color{X: 1.0, Y: 1.0, Z: 1.0}
		blue  = vec.Color{X: 0.5, Y: 0.7, Z: 1.0}
	)
	return white.Scale(1.0 - a).Add(blue.Scale(a))
}

func (center Camera) DefocusDiskSample(defocusDiskU vec.Vec3, defocusDiskV vec.Vec3) vec.Point3 {
	p := vec.RandomInUnitDisk()
	return center.center.Add(defocusDiskU.Scale(p.X)).Add(defocusDiskV.Scale(p.Y))
}
