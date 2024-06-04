package shapes

import (
	"encoding/json"
	"math"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Circle struct {
	color  *value_objects.Color
	center *value_objects.Point
	radius float64
}

func BuildCircle(radius int, color *value_objects.Color, center *value_objects.Point) *Circle {
	return &Circle{
		color:  color,
		center: center,
		radius: float64(radius),
	}
}

func BuildCircleFromJson(objmap map[string]json.RawMessage) *Circle {
	var radius int

	err := json.Unmarshal(objmap["radius"], &radius)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildCircle(radius, color, center)
}

func (c *Circle) Color() *value_objects.Color {
	return c.color
}

func (c *Circle) Intersect(p *value_objects.Point) bool {
	distance := math.Sqrt(float64(((p.X() - c.center.X()) * (p.X() - c.center.X())) + ((p.Y() - c.center.Y()) * (p.Y() - c.center.Y()))))
	return distance <= c.radius
}
