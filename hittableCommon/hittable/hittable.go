package hittable

import (
	"github.com/sergeev-s/raytracer/interval"
	"github.com/sergeev-s/raytracer/vec"
	"github.com/sergeev-s/raytracer/ray"
)

type HitRecord struct {
	P        vec.Point3
	Normal   vec.Point3
	T        float64
	Material Material
	IsFrontFace bool
}

type Material interface {
	Scatter(rayIn ray.Ray, hitRecord HitRecord) (ray.Ray, vec.Color, bool)
}

type Hittable interface {
	Hit(r ray.Ray, rayT interval.Interval) (HitRecord, bool)
}
