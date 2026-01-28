package hittableList

import (
	"github.com/sergeev-s/raytracer/hittable"
	"github.com/sergeev-s/raytracer/interval"
	"github.com/sergeev-s/raytracer/ray"
)

type HittableList struct {
	Hittables []hittable.Hittable
}

func (hittableList *HittableList) Add(hittable hittable.Hittable) {
	hittableList.Hittables = append(hittableList.Hittables, hittable)
}

func (hittableList HittableList) Hit(r ray.Ray, rayT interval.Interval) (hittable.HitRecord, bool) {
	closestHitRecord := hittable.HitRecord{}
	closestT := rayT.Max
	for _, hittable := range hittableList.Hittables {
		interval := interval.Interval{Min: rayT.Min, Max: closestT}
		hitRecord, localHitRecord := hittable.Hit(r, interval)
		if localHitRecord {
			closestHitRecord = hitRecord
			closestT = closestHitRecord.T
		}
	}

	return closestHitRecord, closestT < rayT.Max
}