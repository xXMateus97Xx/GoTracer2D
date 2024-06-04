package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type BottomSemiCircle struct {
	Circle
}

func BuildBottomSemiCircle(radius int, color *value_objects.Color, center *value_objects.Point) *BottomSemiCircle {
	return &BottomSemiCircle{
		Circle: Circle{
			color:  color,
			center: center,
			radius: float64(radius),
		},
	}
}

func BuildBottomSemiCircleFromJson(objmap map[string]json.RawMessage) *BottomSemiCircle {
	var radius int

	err := json.Unmarshal(objmap["radius"], &radius)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildBottomSemiCircle(radius, color, center)
}

func (t *BottomSemiCircle) Intersect(p *value_objects.Point) bool {
	if p.Y() < t.center.Y() {
		return false
	}

	return t.Circle.Intersect(p)
}
