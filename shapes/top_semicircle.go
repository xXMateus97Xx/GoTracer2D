package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type TopSemiCircle struct {
	Circle
}

func BuildTopSemiCircle(radius int, color *value_objects.Color, center *value_objects.Point) *TopSemiCircle {
	return &TopSemiCircle{
		Circle: Circle{
			color:  color,
			center: center,
			radius: float64(radius),
		},
	}
}

func BuildTopSemiCircleFromJson(objmap map[string]json.RawMessage) *TopSemiCircle {
	var radius int

	err := json.Unmarshal(objmap["radius"], &radius)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildTopSemiCircle(radius, color, center)
}

func (t *TopSemiCircle) Intersect(p *value_objects.Point) bool {
	if p.Y() > t.center.Y() {
		return false
	}

	return t.Circle.Intersect(p)
}
