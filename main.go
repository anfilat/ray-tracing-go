package main

import (
	"github.com/anfilat/ray-tracing-go.git/camera"
	"github.com/anfilat/ray-tracing-go.git/hitTable"
	"github.com/anfilat/ray-tracing-go.git/point"
	"github.com/anfilat/ray-tracing-go.git/sphere"
)

func main() {
	world := hitTable.NewList()
	world.Add(sphere.New(point.NewXYZ(0, 0, -1), 0.5))
	world.Add(sphere.New(point.NewXYZ(0, -100.5, -1), 100))

	cam := camera.New()
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400

	cam.Render(world)
}
