package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type RightSemiCircle struct {
	Circle
}

func BuildRightSemiCircle(radius int, color *value_objects.Color, center *value_objects.Point) *RightSemiCircle {
	return &RightSemiCircle{
		Circle: Circle{
			color:  color,
			center: center,
			radius: float64(radius),
		},
	}
}

func BuildRightSemiCircleFromJson(objmap map[string]json.RawMessage) *RightSemiCircle {
	var radius int

	err := json.Unmarshal(objmap["radius"], &radius)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildRightSemiCircle(radius, color, center)
}

func (t *RightSemiCircle) Intersect(p *value_objects.Point) bool {
	if p.X() < t.center.X() {
		return false
	}

	return t.Circle.Intersect(p)
}
