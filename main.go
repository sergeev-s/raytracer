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

// func run() {
// 	materialGround := material.NewLambertian(vec.Color{X: 0.8, Y: 0.8, Z: 0.0})
// 	materialCenter := material.NewLambertian(vec.Color{X: 0.1, Y: 0.2, Z: 0.5})
// 	materialLeftOuter := material.NewDielectric(1.5)
// 	materialLeftInner := material.NewDielectric(1.0 / 1.5)
// 	materialRight := material.NewMetal(vec.Color{X: 0.8, Y: 0.6, Z: 0.2}, 1.0)

// 	world := &hittableList.HittableList{}

// 	sphere1 := sphere.NewSphere(vec.Point3{X: 0, Y: -100.5, Z: -1}, 100, materialGround)
// 	world.Add(&sphere1)

// 	sphere2 := sphere.NewSphere(vec.Point3{X: 0, Y: 0, Z: -1.2}, 0.5, materialCenter)
// 	world.Add(&sphere2)

// 	sphere3 := sphere.NewSphere(vec.Point3{X: -1, Y: 0, Z: -1}, 0.5, materialLeftOuter)
// 	world.Add(&sphere3)

// 	sphere4 := sphere.NewSphere(vec.Point3{X: -1, Y: 0, Z: -1}, 0.4, materialLeftInner)
// 	world.Add(&sphere4)

// 	sphere5 := sphere.NewSphere(vec.Point3{X: 1, Y: 0, Z: -1}, 0.5, materialRight)
// 	world.Add(&sphere5)

// 	cameraInstance := camera.NewCamera(ASPECT_RATIO, IMAGE_WIDTH)
// 	cameraInstance.Render(world)
// }

func run() {
	materialGround := material.NewLambertian(vec.Color{X: 0.8, Y: 0.8, Z: 0.0})

	mirrorLeft := material.NewMetal(vec.Color{X: 0.99, Y: 0.99, Z: 0.99}, 0.0)
	mirrorRight := material.NewMetal(vec.Color{X: 0.95, Y: 0.97, Z: 0.99}, 0.0)

	accent := material.NewLambertian(vec.Color{X: 0.90, Y: 0.25, Z: 0.20})

	world := &hittableList.HittableList{}

	ground := sphere.NewSphere(vec.Point3{X: 0, Y: -100.5, Z: -1}, 100, materialGround)
	world.Add(&ground)

	leftMirrorSphere := sphere.NewSphere(vec.Point3{X: -0.72, Y: 0.10, Z: -2.10}, 0.60, mirrorLeft)
	world.Add(&leftMirrorSphere)

	rightMirrorSphere := sphere.NewSphere(vec.Point3{X: 0.72, Y: 0.10, Z: -2.10}, 0.60, mirrorRight)
	world.Add(&rightMirrorSphere)

	topAccentSphere := sphere.NewSphere(vec.Point3{X: 0.00, Y: 0.95, Z: -2.55}, 0.35, accent)
	world.Add(&topAccentSphere)

	cameraInstance := camera.NewCamera(ASPECT_RATIO, IMAGE_WIDTH, 90, vec.Point3{X: 0, Y: 0, Z: 0}, vec.Point3{X: 0, Y: 0, Z: -1}, vec.Vec3{X: 0, Y: 1, Z: 0})
	cameraInstance.Render(world)
}
