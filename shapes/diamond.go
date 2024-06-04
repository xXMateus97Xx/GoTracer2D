package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Diamond struct {
	color *value_objects.Color
	a, b  *Triangle
}

func BuildDiamond(width int, height int, color *value_objects.Color, center *value_objects.Point) *Diamond {
	halfWidth := width / 2
	halfHeight := height / 2
	left := value_objects.BuildPoint(center.X()-halfWidth, center.Y())
	right := value_objects.BuildPoint(center.X()+halfWidth, center.Y())
	top := value_objects.BuildPoint(center.X(), center.Y()-halfHeight)
	bottom := value_objects.BuildPoint(center.X(), center.Y()+halfHeight)

	return &Diamond{
		color: color,
		a:     BuildTriangle(left, right, top, color),
		b:     BuildTriangle(left, right, bottom, color),
	}
}

func BuildDiamondFromJson(objmap map[string]json.RawMessage) *Diamond {
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

	return BuildDiamond(width, height, color, center)
}

func (d *Diamond) Color() *value_objects.Color {
	return d.color
}

func (d *Diamond) Intersect(p *value_objects.Point) bool {
	return d.a.Intersect(p) || d.b.Intersect(p)
}
