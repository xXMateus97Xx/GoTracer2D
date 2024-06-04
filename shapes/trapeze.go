package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Trapeze struct {
	color       *value_objects.Color
	s           *Square
	left, right *Triangle
}

func BuildTrapeze(height int, topWidth int, bottomWidth int, color *value_objects.Color, center *value_objects.Point) *Trapeze {
	var halfHeight = height / 2
	var leftX = center.X() - topWidth/2
	var rightX = center.X() + topWidth/2
	var topBottomDiff = (bottomWidth - topWidth) / 2
	var topY = center.Y() - halfHeight
	var bottomY = center.Y() + halfHeight

	var topLeft = value_objects.BuildPoint(leftX, topY)
	var bottomLeft = value_objects.BuildPoint(leftX, bottomY)
	var deepestLeft = value_objects.BuildPoint(leftX-topBottomDiff, bottomLeft.Y())

	var topRight = value_objects.BuildPoint(rightX, topY)
	var bottomRight = value_objects.BuildPoint(rightX, bottomY)
	var deepestRight = value_objects.BuildPoint(rightX+topBottomDiff, bottomRight.Y())

	return &Trapeze{
		color: color,
		s:     BuildSquare(topWidth, height, color, center),
		left:  BuildTriangle(topLeft, bottomLeft, deepestLeft, color),
		right: BuildTriangle(topRight, bottomRight, deepestRight, color),
	}
}

func BuildTrapezeFromJson(objmap map[string]json.RawMessage) *Trapeze {
	var height, topWidth, bottomWidth int

	err := json.Unmarshal(objmap["height"], &height)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["topwidth"], &topWidth)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["bottomwidth"], &bottomWidth)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildTrapeze(height, topWidth, bottomWidth, color, center)
}

func (t *Trapeze) Color() *value_objects.Color {
	return t.color
}

func (t *Trapeze) Intersect(p *value_objects.Point) bool {
	return t.s.Intersect(p) || t.left.Intersect(p) || t.right.Intersect(p)
}
