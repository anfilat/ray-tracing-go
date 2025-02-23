package main

import (
	"github.com/anfilat/ray-tracing-go.git/camera"
	"github.com/anfilat/ray-tracing-go.git/color"
	"github.com/anfilat/ray-tracing-go.git/list"
	"github.com/anfilat/ray-tracing-go.git/material"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/sphere"
	"math"
)

func main() {
	world := list.New()

	r := math.Cos(math.Pi / 4)

	materialLeft := material.NewLambertian(color.NewRGB(0, 0, 1))
	materialRight := material.NewLambertian(color.NewRGB(1, 0, 0))

	world.Add(sphere.New(point.NewXYZ(-r, 0, -1), r, materialLeft))
	world.Add(sphere.New(point.NewXYZ(r, 0, -1), r, materialRight))

	cam := camera.New()
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400
	cam.SamplesPerPixel = 100
	cam.MaxDepth = 50
	cam.Vfov = 90

	cam.Render(world)
}
