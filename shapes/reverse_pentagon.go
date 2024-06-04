package shapes

import (
	"encoding/json"
	"math"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type ReversePentagon struct {
	color   *value_objects.Color
	a, b, c *Triangle
}

func BuildReversePentagon(side int, color *value_objects.Color, center *value_objects.Point) *ReversePentagon {
	sideFloat := float64(side)
	height := sideFloat * 3.0776834 / 2
	halfHeight := height / 2
	diagonal := sideFloat * 3.236068 / 2
	halfDiagonal := diagonal / 2
	halfSide := sideFloat / 2

	centerYFloat := float64(center.Y())
	centerXFloat := float64(center.X())
	bottomY := centerYFloat + halfHeight
	bottom := value_objects.BuildPoint(center.X(), int(bottomY))
	topLeft := value_objects.BuildPoint(int(centerXFloat-halfSide), int(centerYFloat-halfHeight))
	topRight := value_objects.BuildPoint(int(centerXFloat+halfSide), int(centerYFloat-halfHeight))
	sideCornerY := (int)(bottomY - math.Sqrt((sideFloat*sideFloat)-(halfDiagonal*halfDiagonal)))

	return &ReversePentagon{
		color: color,
		a:     BuildTriangle(bottom, topLeft, topRight, color),
		b: BuildTriangle(bottom, topLeft,
			value_objects.BuildPoint(int(centerXFloat-halfDiagonal), sideCornerY),
			color),
		c: BuildTriangle(bottom, topRight,
			value_objects.BuildPoint(int(centerXFloat+halfDiagonal), sideCornerY),
			color),
	}
}

func BuildReversePentagonFromJson(objmap map[string]json.RawMessage) *ReversePentagon {
	var side int

	err := json.Unmarshal(objmap["side"], &side)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildReversePentagon(side, color, center)
}

func (p *ReversePentagon) Color() *value_objects.Color {
	return p.color
}

func (p *ReversePentagon) Intersect(point *value_objects.Point) bool {
	return p.a.Intersect(point) || p.b.Intersect(point) || p.c.Intersect((point))
}
