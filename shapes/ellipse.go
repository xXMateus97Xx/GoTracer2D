package shapes

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Ellipse struct {
	color          *value_objects.Color
	center, radius *value_objects.Point
}

func BuildEllipse(radius *value_objects.Point, color *value_objects.Color, center *value_objects.Point) *Ellipse {
	return &Ellipse{
		color:  color,
		center: center,
		radius: radius,
	}
}

func BuildEllipseFromJson(objmap map[string]json.RawMessage) *Ellipse {
	color := value_objects.BuildColorFromJson(objmap["color"])
	center := value_objects.BuildPointFromJson(objmap["center"])
	radius := value_objects.BuildPointFromJson(objmap["radius"])

	return BuildEllipse(radius, color, center)
}

func (e *Ellipse) Color() *value_objects.Color {
	return e.color
}

func (e *Ellipse) Intersect(p *value_objects.Point) bool {
	return ((p.X()-e.center.X())*(p.X()-e.center.X())/e.radius.X())+((p.Y()-e.center.Y())*(p.Y()-e.center.Y())/e.radius.Y()) <= e.radius.X()
}
