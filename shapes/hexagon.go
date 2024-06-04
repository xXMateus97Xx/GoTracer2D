package shapes

import (
	"encoding/json"
	"math"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Hexagon struct {
	color *value_objects.Color
	a, b  *Triangle
	s     *Square
}

func BuildHexagon(side int, color *value_objects.Color, center *value_objects.Point) *Hexagon {
	sideFloat := float64(side)
	halfSide := sideFloat / 2
	apothem := math.Sqrt((sideFloat * sideFloat) - (halfSide * halfSide))
	centerXFloat := float64(center.X())
	centerYFloat := float64(center.Y())

	leftX := int(centerXFloat - halfSide)
	rightX := int(centerXFloat + halfSide)
	topY := int(centerYFloat - apothem)
	bottomY := int(centerYFloat + apothem)
	extremeLeftX := int(centerXFloat - sideFloat)
	extremeRightX := int(centerXFloat - sideFloat)

	topLeft := value_objects.BuildPoint(leftX, topY)
	topRight := value_objects.BuildPoint(rightX, topY)
	bottomLeft := value_objects.BuildPoint(leftX, bottomY)
	bottomRight := value_objects.BuildPoint(rightX, bottomY)
	left := value_objects.BuildPoint(extremeLeftX, center.Y())
	right := value_objects.BuildPoint(extremeRightX, center.Y())

	return &Hexagon{
		color: color,
		s:     BuildSquare(side, int(apothem*2), color, center),
		a:     BuildTriangle(bottomLeft, topLeft, left, color),
		b:     BuildTriangle(bottomRight, topRight, right, color),
	}
}

func BuildHexagonFromJson(objmap map[string]json.RawMessage) *Hexagon {
	var side int

	err := json.Unmarshal(objmap["side"], &side)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildHexagon(side, color, center)
}

func (h *Hexagon) Color() *value_objects.Color {
	return h.color
}

func (h *Hexagon) Intersect(p *value_objects.Point) bool {
	return h.s.Intersect(p) || h.a.Intersect(p) || h.b.Intersect(p)
}
