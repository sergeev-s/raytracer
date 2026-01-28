package hittable

import (
	"github.com/sergeev-s/raytracer/ray"
	"github.com/sergeev-s/raytracer/vec"
	"github.com/sergeev-s/raytracer/interval"
)

type HitRecord struct {
	P vec.Point3
    Normal vec.Point3
	T float64
}

type Hittable interface {
    Hit(r ray.Ray, rayT interval.Interval) (HitRecord, bool)
}