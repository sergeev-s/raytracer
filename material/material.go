package material

import (
	"github.com/sergeev-s/raytracer/hittableCommon/hittable"
	"github.com/sergeev-s/raytracer/ray"
	"github.com/sergeev-s/raytracer/vec"
)

type Lambertian struct {
	Albedo vec.Color
}

type Metal struct {
	Albedo vec.Color
}

func NewLambertian(albedo vec.Color) Lambertian {
	return Lambertian{Albedo: albedo}
}

func NewMetal(albedo vec.Color) Metal {
	return Metal{Albedo: albedo}
}

func (material Lambertian) Scatter(rayIn ray.Ray, hitRecord hittable.HitRecord) (ray.Ray, vec.Color, bool) {
	// _ = rayIn
	scatterDirection := hitRecord.Normal.Add(vec.RandomUnitVector())
	if scatterDirection.NearZero() {
		scatterDirection = hitRecord.Normal
	}

	scatteredRay := ray.Ray{Origin: hitRecord.P, Direction: scatterDirection}
	return scatteredRay, material.Albedo, true
}

func (material Metal) Scatter(rayIn ray.Ray, hitRecord hittable.HitRecord) (ray.Ray, vec.Color, bool) {
	reflected := vec.Reflect(rayIn.Direction, hitRecord.Normal)
	reflectedRay := ray.Ray{Origin: hitRecord.P, Direction: reflected}
	return reflectedRay, material.Albedo, reflected.Dot(hitRecord.Normal) > 0
}	
