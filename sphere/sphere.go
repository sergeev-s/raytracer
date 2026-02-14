package sphere

import (
	"math"

	"github.com/sergeev-s/raytracer/hittableCommon/hittable"
	"github.com/sergeev-s/raytracer/interval"
	"github.com/sergeev-s/raytracer/ray"
	"github.com/sergeev-s/raytracer/vec"
)

type Sphere struct {
	Center   vec.Point3
	Radius   float64
	Material hittable.Material
}

func NewSphere(center vec.Point3, radius float64, material hittable.Material) Sphere {
	return Sphere{
		Center:   center,
		Radius:   math.Max(0, radius),
		Material: material,
	}
}

func (sphere Sphere) Hit(r ray.Ray, rayT interval.Interval) (hittable.HitRecord, bool) {
	oc := sphere.Center.Sub(r.Origin)
	a := r.Direction.LengthSquared()
	b := -2 * r.Direction.Dot(oc)
	c := oc.LengthSquared() - sphere.Radius*sphere.Radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return hittable.HitRecord{}, false
	}

	root := (-b - math.Sqrt(discriminant)) / (2 * a)

	if !rayT.Surrounds(root) {
		root = (-b + math.Sqrt(discriminant)) / (2 * a)
		if !rayT.Surrounds(root) {
			return hittable.HitRecord{}, false
		}
	}

	p := r.At(root)
	outwardNormal := p.Sub(sphere.Center).Divide(sphere.Radius)
	isFrontFace := r.Direction.Dot(outwardNormal) < 0
	normal := outwardNormal
	if !isFrontFace {
		normal = outwardNormal.Negate()
	}

	return hittable.HitRecord{
		P:        p,
		Normal:   normal,
		T:        root,
		Material: sphere.Material,
		IsFrontFace: isFrontFace,
	}, true
}
