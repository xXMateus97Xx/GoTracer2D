package shapes

import (
	"encoding/json"
	"math"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Ring struct {
	color             *value_objects.Color
	center            *value_objects.Point
	radius, thickness float64
}

func BuildRing(radius int, thickness int, color *value_objects.Color, center *value_objects.Point) *Ring {
	return &Ring{
		color:     color,
		center:    center,
		radius:    float64(radius),
		thickness: float64(thickness),
	}
}

func BuildRingFromJson(objmap map[string]json.RawMessage) *Ring {
	var radius, thickness int

	err := json.Unmarshal(objmap["radius"], &radius)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["thickness"], &thickness)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildRing(radius, thickness, color, center)
}

func (r *Ring) Color() *value_objects.Color {
	return r.color
}

func (r *Ring) Intersect(p *value_objects.Point) bool {
	distance := math.Sqrt(float64(((p.X() - r.center.X()) * (p.X() - r.center.X())) + ((p.Y() - r.center.Y()) * (p.Y() - r.center.Y()))))
	return distance <= r.radius && distance > r.thickness
}
