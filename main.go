package main

import (
	"github.com/sergeev-s/raytracer/hittable/list"
	"github.com/sergeev-s/raytracer/sphere"
	"github.com/sergeev-s/raytracer/vec"
	"github.com/sergeev-s/raytracer/camera"
)

const (
	IMAGE_WIDTH     = 400
	ASPECT_RATIO    = 16.0 / 9.0
)

func main() {
	run()
}

func run() {
	world := &hittableList.HittableList{}
	
	sphere1 := sphere.NewSphere(vec.Point3{X: 0, Y: 0, Z: -1}, 0.5)
	world.Add(&sphere1)
	
	sphere2 := sphere.NewSphere(vec.Point3{X: 0, Y: -100.5, Z: -1}, 100)
	world.Add(&sphere2)		
	
	cameraInstance := camera.NewCamera(ASPECT_RATIO, IMAGE_WIDTH)
	cameraInstance.Render(world)
}
