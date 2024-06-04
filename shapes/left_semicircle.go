package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type LeftSemiCircle struct {
	Circle
}

func BuildLeftSemiCircle(radius int, color *value_objects.Color, center *value_objects.Point) *LeftSemiCircle {
	return &LeftSemiCircle{
		Circle: Circle{
			color:  color,
			center: center,
			radius: float64(radius),
		},
	}
}

func BuildLeftSemiCircleFromJson(objmap map[string]json.RawMessage) *LeftSemiCircle {
	var radius int

	err := json.Unmarshal(objmap["radius"], &radius)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildLeftSemiCircle(radius, color, center)
}

func (t *LeftSemiCircle) Intersect(p *value_objects.Point) bool {
	if p.X() > t.center.X() {
		return false
	}

	return t.Circle.Intersect(p)
}
