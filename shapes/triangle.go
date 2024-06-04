package shapes

import (
	"encoding/json"
	"math"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Triangle struct {
	color      *value_objects.Color
	v0, v1, v2 *value_objects.Point
	area       float64
}

func BuildTriangle(v0 *value_objects.Point, v1 *value_objects.Point, v2 *value_objects.Point, color *value_objects.Color) *Triangle {
	return &Triangle{
		color: color,
		v0:    v0,
		v1:    v1,
		v2:    v2,
		area:  triangleArea(v0, v1, v2),
	}
}

func BuildTriangleFromJson(objmap map[string]json.RawMessage) *Triangle {
	color := value_objects.BuildColorFromJson(objmap["color"])
	v0 := value_objects.BuildPointFromJson(objmap["v0"])
	v1 := value_objects.BuildPointFromJson(objmap["v1"])
	v2 := value_objects.BuildPointFromJson(objmap["v2"])

	return BuildTriangle(v0, v1, v2, color)
}

func (t *Triangle) Color() *value_objects.Color {
	return t.color
}

func (t *Triangle) Intersect(p *value_objects.Point) bool {
	var a1 = triangleArea(p, t.v1, t.v2)
	var a2 = triangleArea(t.v0, p, t.v2)
	var a3 = triangleArea(t.v0, t.v1, p)

	return t.area == a1+a2+a3
}

func triangleArea(p1 *value_objects.Point, p2 *value_objects.Point, p3 *value_objects.Point) float64 {
	return math.Abs(float64(p1.X()*(p2.Y()-p3.Y())+p2.X()*(p3.Y()-p1.Y())+p3.X()*(p1.Y()-p2.Y())) / 2.0)
}
