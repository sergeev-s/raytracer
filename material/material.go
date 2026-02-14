package material

import (
	"math"
	"math/rand"

	"github.com/sergeev-s/raytracer/hittableCommon/hittable"
	"github.com/sergeev-s/raytracer/ray"
	"github.com/sergeev-s/raytracer/vec"
)

type Lambertian struct {
	Albedo vec.Color
}

type Metal struct {
	Albedo vec.Color
	Fuzz   float64
}

type Dielectric struct {
	RefractionIndex float64
}

func NewLambertian(albedo vec.Color) Lambertian {
	return Lambertian{Albedo: albedo}
}

func NewMetal(albedo vec.Color, fuzz float64) Metal {
	if fuzz > 1 {
		fuzz = 1.0
	}
	return Metal{Albedo: albedo, Fuzz: fuzz}
}

func NewDielectric(refractionIndex float64) Dielectric {
	return Dielectric{RefractionIndex: refractionIndex}
}

func (material Lambertian) Scatter(rayIn ray.Ray, hitRecord hittable.HitRecord) (ray.Ray, vec.Color, bool) {
	scatterDirection := hitRecord.Normal.Add(vec.RandomUnitVector())
	if scatterDirection.NearZero() {
		scatterDirection = hitRecord.Normal
	}

	scatteredRay := ray.Ray{Origin: hitRecord.P, Direction: scatterDirection}
	return scatteredRay, material.Albedo, true
}

func (material Metal) Scatter(rayIn ray.Ray, hitRecord hittable.HitRecord) (ray.Ray, vec.Color, bool) {
	reflected := vec.Reflect(rayIn.Direction, hitRecord.Normal)
	reflectedUnit := reflected.GetUnitVec()
	fuzzVector := vec.RandomUnitVector().Scale(material.Fuzz)
	totalReflected := reflectedUnit.Add(fuzzVector)

	if totalReflected.Dot(hitRecord.Normal) < 0 {
		return ray.Ray{}, vec.Color{}, false
	}

	reflectedRay := ray.Ray{Origin: hitRecord.P, Direction: reflectedUnit.Add(fuzzVector)}
	return reflectedRay, material.Albedo, true
}

func (material Dielectric) Scatter(rayIn ray.Ray, hitRecord hittable.HitRecord) (ray.Ray, vec.Color, bool) {
	var totalRI float64

	if hitRecord.IsFrontFace == true {
		totalRI = 1.0 / material.RefractionIndex
	} else {
		totalRI = material.RefractionIndex
	}

	rInUnit := rayIn.Direction.GetUnitVec()
	// refracted := rInUnit.Refract(hitRecord.Normal, totalRI)

	// reflectedRay := ray.Ray{Origin: hitRecord.P, Direction: refracted}

	cosTheta := math.Min(rInUnit.Negate().Dot(hitRecord.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := totalRI*sinTheta > 1.0

	var finalDirection vec.Vec3

	if cannotRefract || material.Reflectance(cosTheta, totalRI) > rand.Float64() {
		finalDirection = vec.Reflect(rInUnit, hitRecord.Normal)
	} else {
		finalDirection = rInUnit.Refract(hitRecord.Normal, totalRI)
	}

	reflactedRay := ray.Ray{Origin: hitRecord.P, Direction: finalDirection}

	return reflactedRay, vec.Color{X: 1.0, Y: 1.0, Z: 1.0}, true
}

func (material Dielectric) Reflectance(cosine, totalRI float64) float64 {
	r0 := (1 - totalRI) / (1 + totalRI) * (1 - totalRI) / (1 + totalRI)
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
