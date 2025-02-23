package main

import (
	"github.com/anfilat/ray-tracing-go.git/camera"
	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/common"
	"github.com/anfilat/ray-tracing-go.git/list"
	"github.com/anfilat/ray-tracing-go.git/material"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/sphere"
	"github.com/anfilat/ray-tracing-go.git/vec3"
)

func main() {
	world := list.New()

	groundMaterial := material.NewLambertian(color.NewRGB(0.5, 0.5, 0.5))
	world.Add(sphere.New(point.NewXYZ(0, -1000, 0), 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := common.Random()
			center := point.NewXYZ(float64(a)+0.9*common.Random(), 0.2, float64(b)+0.9*common.Random())

			if center.Sub(point.NewXYZ(4, 0.2, 0)).Length() > 0.9 {
				if chooseMat < 0.8 {
					// diffuse
					albedo := vec3.Random().Mul(vec3.Random())
					sphereMaterial := material.NewLambertian(albedo)
					world.Add(sphere.New(center, 0.2, sphereMaterial))
				} else if chooseMat < 0.95 {
					// metal
					albedo := vec3.RandomMM(0.5, 1)
					fuzz := common.RandomMM(0, 0.5)
					sphereMaterial := material.NewMetal(albedo, fuzz)
					world.Add(sphere.New(center, 0.2, sphereMaterial))
				} else {
					// glass
					sphereMaterial := material.NewDielectric(1.5)
					world.Add(sphere.New(center, 0.2, sphereMaterial))
				}
			}
		}
	}

	material1 := material.NewDielectric(1.5)
	world.Add(sphere.New(point.NewXYZ(0, 1, 0), 1, material1))

	material2 := material.NewLambertian(color.NewRGB(0.4, 0.2, 0.1))
	world.Add(sphere.New(point.NewXYZ(-4, 1, 0), 1, material2))

	material3 := material.NewMetal(color.NewRGB(0.7, 0.6, 0.5), 0)
	world.Add(sphere.New(point.NewXYZ(4, 1, 0), 1, material3))

	cam := camera.New()
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 1200
	cam.SamplesPerPixel = 500
	cam.MaxDepth = 50
	cam.Vfov = 20
	cam.LookFrom = point.NewXYZ(13, 2, 3)
	cam.LookAt = point.NewXYZ(0, 0, 0)
	cam.Vup = vec3.New(0, 1, 0)
	cam.DefocusAngle = 0.6
	cam.FocusDist = 10

	cam.Render(world)
}
