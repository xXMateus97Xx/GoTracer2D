package tracer

import (
	"encoding/json"

	"github.com/xXMateus97Xx/gotracer2d/shapes"
	"github.com/xXMateus97Xx/gotracer2d/value_objects"
)

type Shape interface {
	Color() *value_objects.Color
	Intersect(p *value_objects.Point) bool
}

type ShapeType int

const (
	SquareType ShapeType = iota
	CircleType
	TriangleType
	DiamondType
	PentagonType
	HexagonType
	ReversePentagonType
	EllipseType
	TrapezeType
	ReverseTrapezeType
	RingType
	TopSemiCircleType
	RightSemiCircleType
	BottomSemiCircleType
	LeftSemiCircleType
)

func BuildShapesFromJson(j json.RawMessage) []Shape {
	var mapSlice []map[string]json.RawMessage
	err := json.Unmarshal(j, &mapSlice)

	if err != nil {
		panic(err)
	}

	shapesArr := make([]Shape, len(mapSlice))

	for i := 0; i < len(mapSlice); i++ {
		shapeObj := mapSlice[i]

		var sType ShapeType
		err = json.Unmarshal(shapeObj["type"], &sType)

		if err != nil {
			panic(err)
		}

		shapesArr[i] = buildShapeFromType(sType, shapeObj)
	}

	return shapesArr
}

func buildShapeFromType(sType ShapeType, objmap map[string]json.RawMessage) Shape {
	switch sType {
	case SquareType:
		return shapes.BuildSquareFromJson(objmap)
	case CircleType:
		return shapes.BuildCircleFromJson(objmap)
	case TriangleType:
		return shapes.BuildTriangleFromJson(objmap)
	case DiamondType:
		return shapes.BuildDiamondFromJson(objmap)
	case PentagonType:
		return shapes.BuildPentagonFromJson(objmap)
	case HexagonType:
		return shapes.BuildHexagonFromJson(objmap)
	case ReversePentagonType:
		return shapes.BuildReversePentagonFromJson(objmap)
	case EllipseType:
		return shapes.BuildEllipseFromJson(objmap)
	case TrapezeType:
		return shapes.BuildTrapezeFromJson(objmap)
	case ReverseTrapezeType:
		return shapes.BuildReverseTrapezeFromJson(objmap)
	case RingType:
		return shapes.BuildRingFromJson(objmap)
	case TopSemiCircleType:
		return shapes.BuildTopSemiCircleFromJson(objmap)
	case RightSemiCircleType:
		return shapes.BuildRightSemiCircleFromJson(objmap)
	case BottomSemiCircleType:
		return shapes.BuildBottomSemiCircleFromJson(objmap)
	case LeftSemiCircleType:
		return shapes.BuildLeftSemiCircleFromJson(objmap)
	default:
		panic("type not found")
	}
}
