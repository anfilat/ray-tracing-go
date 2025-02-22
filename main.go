package main

import (
	"github.com/anfilat/ray-tracing-go.git/camera"
	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/hitTable"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/sphere"
)

func main() {
	world := hitTable.NewList()

	materialGround := hitTable.NewLambertian(color.NewRGB(0.8, 0.8, 0))
	materialCenter := hitTable.NewLambertian(color.NewRGB(0.1, 0.2, 0.5))
	materialLeft := hitTable.NewDielectric(1.5)
	materialRight := hitTable.NewMetal(color.NewRGB(0.8, 0.6, 0.2), 1)

	world.Add(sphere.New(point.NewXYZ(0, -100.5, -1), 100, materialGround))
	world.Add(sphere.New(point.NewXYZ(0, 0, -1.2), 0.5, materialCenter))
	world.Add(sphere.New(point.NewXYZ(-1, 0, -1), 0.5, materialLeft))
	world.Add(sphere.New(point.NewXYZ(1, 0, -1), 0.5, materialRight))

	cam := camera.New()
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400
	cam.SamplesPerPixel = 100
	cam.MaxDepth = 50

	cam.Render(world)
}
