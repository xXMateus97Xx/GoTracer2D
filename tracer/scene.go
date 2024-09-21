package tracer

import (
	"encoding/json"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Scene struct {
	background    *value_objects.Color
	shapes        []Shape
	width, height int
}

func BuildScene(background *value_objects.Color, shapes []Shape, width int, height int) *Scene {
	return &Scene{
		background: background,
		shapes:     shapes,
		width:      width,
		height:     height,
	}
}

func BuildSceneFromJson(path string) *Scene {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var objmap map[string]json.RawMessage
	err = json.Unmarshal(bytes, &objmap)

	if err != nil {
		panic(err)
	}

	background := value_objects.BuildColorFromJson(objmap["background"])

	var width, height int
	err = json.Unmarshal(objmap["width"], &width)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["height"], &height)
	if err != nil {
		panic(err)
	}

	shapes := BuildShapesFromJson(objmap["shapes"])

	return BuildScene(background, shapes, width, height)
}

func (s *Scene) Render(outputPath string) {
	file, err := os.Create(outputPath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s.RenderFile(file)
}

func (s *Scene) RenderFile(file *os.File) {

	const bufferSize = 1023 * 4

	channCapacity := int32(math.Ceil(float64(s.height*s.width*3)/bufferSize)) + 1

	chann := make(chan []byte, channCapacity)
	done := make(chan bool)

	go writeFile(chann, done, file)

	s.writeHeader(chann)

	p := value_objects.BuildPoint(0, 0)

	bytes := make([]byte, bufferSize)
	pos := 0

	for y := 0; y < s.height; y++ {
		p.SetY(y)

		for x := 0; x < s.width; x++ {
			p.SetX(x)

			color := s.background

			for i := len(s.shapes) - 1; i >= 0; i-- {
				s := s.shapes[i]

				if s.Intersect(p) {
					color = color.Add(s.Color())
					break
				}
			}

			color.ToBytes(bytes[pos : pos+3])
			pos += 3

			if pos == len(bytes) {
				chann <- bytes
				bytes = make([]byte, bufferSize)
				pos = 0
			}
		}
	}

	if pos > 0 {
		chann <- bytes[0:pos]
	}

	close(chann)

	<-done
}

func writeFile(chann chan []byte, done chan bool, file *os.File) {
	for b := range chann {
		file.Write(b)
	}

	done <- true
}

func (s *Scene) writeHeader(chann chan []byte) {
	h1 := []byte("P6\n")
	h2 := []byte(" ")
	h3 := []byte("\n255\n")

	headerLength := len(h1) + len(h2) + len(h3) + 20

	bytes := make([]byte, headerLength)

	wrote := copy(bytes, h1)
	wrote += copy(bytes[wrote:], []byte(strconv.Itoa(s.width)))
	wrote += copy(bytes[wrote:], h2)
	wrote += copy(bytes[wrote:], []byte(strconv.Itoa(s.height)))
	wrote += copy(bytes[wrote:], h3)

	chann <- bytes[0:wrote]
}
