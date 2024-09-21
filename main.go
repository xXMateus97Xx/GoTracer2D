package main

import (
	"fmt"
	"os"

	"github.com/xXMateus97Xx/gotracer2d/tracer"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("gotracer2d input.json output.ppm")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	scene := tracer.BuildSceneFromJson(input)

	scene.Render(output)
}
