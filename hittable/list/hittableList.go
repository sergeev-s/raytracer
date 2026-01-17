package hittableList

import (
	"github.com/sergeev-s/raytracer/hittable"
	"github.com/sergeev-s/raytracer/ray"
)

type HittableList struct {
	Hittables []hittable.Hittable
}

func (hittableList *HittableList) Add(hittable hittable.Hittable) {
	hittableList.Hittables = append(hittableList.Hittables, hittable)
}

func (hittableList *HittableList) Hit(r ray.Ray, rayTMin float64, rayTMax float64) (*hittable.HitRecord, bool) {
	closestHitRecord := hittable.HitRecord{}
	closestT := rayTMax
	for _, hittable := range hittableList.Hittables {
		hitRecord, localHitRecord := hittable.Hit(r, rayTMin, closestT)
		if localHitRecord {
			closestHitRecord = *hitRecord
			closestT = closestHitRecord.T
		}
	}

	return &closestHitRecord, closestT < rayTMax
}