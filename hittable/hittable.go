package hittable

import (
	"github.com/sergeev-s/raytracer/ray"
	"github.com/sergeev-s/raytracer/vec"
)

type HitRecord struct {
	P vec.Point3
    Normal vec.Point3
	T float64
}

type Hittable interface {
    Hit(r ray.Ray, tMin, tMax float64) (*HitRecord, bool)
}