package shapes

import (
	"encoding/json"
	"math"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Pentagon struct {
	color   *value_objects.Color
	a, b, c *Triangle
}

func BuildPentagon(side int, color *value_objects.Color, center *value_objects.Point) *Pentagon {
	sideFloat := float64(side)
	height := sideFloat * 3.0776834 / 2
	halfHeight := height / 2
	diagonal := sideFloat * 3.236068 / 2
	halfDiagonal := diagonal / 2
	halfSide := sideFloat / 2

	centerYFloat := float64(center.Y())
	centerXFloat := float64(center.X())
	var top = value_objects.BuildPoint(center.X(), int(centerYFloat-halfHeight))
	var baseLeft = value_objects.BuildPoint(int(centerXFloat-halfSide), int(centerYFloat+halfHeight))
	var baseRight = value_objects.BuildPoint(int(centerXFloat+halfSide), int(centerYFloat+halfHeight))
	var sideCornerY = int(float64(top.Y()) + math.Sqrt((sideFloat*sideFloat)-(halfDiagonal*halfDiagonal)))

	return &Pentagon{
		color: color,
		a:     BuildTriangle(top, baseLeft, baseRight, color),
		b: BuildTriangle(top, baseLeft,
			value_objects.BuildPoint(int(centerXFloat-halfDiagonal), sideCornerY),
			color),
		c: BuildTriangle(top, baseRight,
			value_objects.BuildPoint(int(centerXFloat+halfDiagonal), sideCornerY),
			color),
	}
}

func BuildPentagonFromJson(objmap map[string]json.RawMessage) *Pentagon {
	var side int

	err := json.Unmarshal(objmap["side"], &side)
	if err != nil {
		panic(err)
	}

	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])

	return BuildPentagon(side, color, center)
}

func (p *Pentagon) Color() *value_objects.Color {
	return p.color
}

func (p *Pentagon) Intersect(point *value_objects.Point) bool {
	return p.a.Intersect(point) || p.b.Intersect(point) || p.c.Intersect((point))
}
