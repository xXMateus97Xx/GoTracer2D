package main

import (
	"os"

	"github.com/xXMateus97Xx/gotracer2d/tracer"
)

func main() {

	input := os.Args[1]
	output := os.Args[2]

	scene := tracer.BuildSceneFromJson(input)

	scene.Render(output)
}
