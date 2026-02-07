package main

import (
	"github.com/sergeev-s/raytracer/camera"
	"github.com/sergeev-s/raytracer/hittableCommon/list"
	"github.com/sergeev-s/raytracer/material"
	"github.com/sergeev-s/raytracer/sphere"
	"github.com/sergeev-s/raytracer/vec"
)

const (
	IMAGE_WIDTH  = 400
	ASPECT_RATIO = 16.0 / 9.0
)

func main() {
	run()
}

func run() {
	materialGround := material.NewLambertian(vec.Color{X: 0.8, Y: 0.8, Z: 0.0})
	materialCenter := material.NewLambertian(vec.Color{X: 0.1, Y: 0.2, Z: 0.5})
	materialLeft := material.NewMetal(vec.Color{X: 0.8, Y: 0.8, Z: 0.8})
	materialRight := material.NewMetal(vec.Color{X: 0.8, Y: 0.6, Z: 0.2})

	world := &hittableList.HittableList{}

	sphere1 := sphere.NewSphere(vec.Point3{X: 0, Y: -100.5, Z: -1}, 100, materialGround)
	world.Add(&sphere1)

	sphere2 := sphere.NewSphere(vec.Point3{X: 0, Y: 0, Z: -1.2}, 0.5, materialCenter)
	world.Add(&sphere2)

	sphere3 := sphere.NewSphere(vec.Point3{X: -1, Y: 0, Z: -1}, 0.5, materialLeft)
	world.Add(&sphere3)

	sphere4 := sphere.NewSphere(vec.Point3{X: 1, Y: 0, Z: -1}, 0.5, materialRight)
	world.Add(&sphere4)	

	cameraInstance := camera.NewCamera(ASPECT_RATIO, IMAGE_WIDTH)
	cameraInstance.Render(world)
}
