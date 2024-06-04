package value_objects

import "encoding/json"

type Point struct {
	x, y int
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) SetX(x int) {
	p.x = x
}

func (p *Point) Y() int {
	return p.y
}

func (p *Point) SetY(y int) {
	p.y = y
}

func BuildPoint(x int, y int) *Point {
	return &Point{x: x, y: y}
}

func BuildPointFromJson(j json.RawMessage) *Point {
	var x, y int

	var objmap map[string]json.RawMessage
	err := json.Unmarshal(j, &objmap)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["x"], &x)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["y"], &y)
	if err != nil {
		panic(err)
	}

	return BuildPoint(x, y)
}
