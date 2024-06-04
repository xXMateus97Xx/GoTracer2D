package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Square struct {
	color                        *value_objects.Color
	topY, bottomY, leftX, rightX int
}

func BuildSquare(width int, height int, color *value_objects.Color, center *value_objects.Point) *Square {
	halfHeight := height / 2
	halfWidth := width / 2

	return &Square{
		color:   color,
		topY:    center.Y() - halfHeight,
		bottomY: center.Y() + halfHeight,
		leftX:   center.X() - halfWidth,
		rightX:  center.X() + halfWidth,
	}
}

func BuildSquareFromJson(objmap map[string]json.RawMessage) *Square {
	var width, height int

	err := json.Unmarshal(objmap["width"], &width)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["height"], &height)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildSquare(width, height, color, center)
}

func (s *Square) Color() *value_objects.Color {
	return s.color
}

func (s *Square) Intersect(p *value_objects.Point) bool {
	return p.X() >= s.leftX && p.X() <= s.rightX && p.Y() >= s.topY && p.Y() <= s.bottomY
}
