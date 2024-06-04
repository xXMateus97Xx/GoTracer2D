package value_objects

import "encoding/json"

type Color struct {
	r, g, b byte
	a       float64
}

func BuildColor(r byte, g byte, b byte) *Color {
	return &Color{r: r, g: g, b: b, a: 1}
}

func BuildAlphaColor(r byte, g byte, b byte, a float64) *Color {
	return &Color{r: r, g: g, b: b, a: a}
}

func BuildColorFromJson(j json.RawMessage) *Color {
	var objmap map[string]json.RawMessage
	err := json.Unmarshal(j, &objmap)

	if err != nil {
		panic(err)
	}

	var r, g, b byte
	var a float64

	err = json.Unmarshal(objmap["r"], &r)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["g"], &g)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(objmap["b"], &b)
	if err != nil {
		panic(err)
	}

	rawA, ok := objmap["a"]
	if ok {
		err = json.Unmarshal(rawA, &a)
		if err != nil {
			panic(err)
		}

		return BuildAlphaColor(r, g, b, a)
	}

	return BuildColor(r, g, b)
}

func (c *Color) R() byte {
	return c.r
}

func (c *Color) G() byte {
	return c.g
}

func (c *Color) B() byte {
	return c.b
}

func (c *Color) A() float64 {
	return c.a
}

func (c *Color) Add(c2 *Color) *Color {
	alpha := c2.a

	return BuildColor(
		byte((1-alpha)*float64(c.r)+alpha*float64(c2.r)),
		byte((1-alpha)*float64(c.g)+alpha*float64(c2.g)),
		byte((1-alpha)*float64(c.b)+alpha*float64(c2.b)))
}

func (c *Color) ToBytes(b []byte) {
	b[0] = c.r
	b[1] = c.g
	b[2] = c.b
}
