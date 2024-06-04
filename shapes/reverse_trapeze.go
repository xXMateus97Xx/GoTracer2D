package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type ReverseTrapeze struct {
	color       *value_objects.Color
	s           *Square
	left, right *Triangle
}

func BuildReverseTrapeze(height int, topWidth int, bottomWidth int, color *value_objects.Color, center *value_objects.Point) *ReverseTrapeze {
	halfHeight := height / 2
	leftX := center.X() - bottomWidth/2
	rightX := center.X() + bottomWidth/2
	topBottomDiff := (topWidth - bottomWidth) / 2
	topY := center.Y() - halfHeight
	bottomY := center.Y() + halfHeight

	topLeft := value_objects.BuildPoint(leftX, topY)
	bottomLeft := value_objects.BuildPoint(leftX, bottomY)
	deepestLeft := value_objects.BuildPoint(leftX-topBottomDiff, topLeft.Y())

	topRight := value_objects.BuildPoint(rightX, topY)
	bottomRight := value_objects.BuildPoint(rightX, bottomY)
	deepestRight := value_objects.BuildPoint(rightX+topBottomDiff, topRight.Y())
	return &ReverseTrapeze{
		color: color,
		s:     BuildSquare(bottomWidth, height, color, center),
		left:  BuildTriangle(topLeft, bottomLeft, deepestLeft, color),
		right: BuildTriangle(topRight, bottomRight, deepestRight, color),
	}
}

func BuildReverseTrapezeFromJson(objmap map[string]json.RawMessage) *ReverseTrapeze {
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

	return BuildReverseTrapeze(height, topWidth, bottomWidth, color, center)
}

func (t *ReverseTrapeze) Color() *value_objects.Color {
	return t.color
}

func (t *ReverseTrapeze) Intersect(p *value_objects.Point) bool {
	return t.s.Intersect(p) || t.left.Intersect(p) || t.right.Intersect(p)
}
