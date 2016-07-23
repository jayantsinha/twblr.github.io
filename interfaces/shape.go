package interfaces

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

type Rectangle struct {
	length int
	breadth int
}

type Hybrid struct {
	s Square
	r Rectangle
}

func (s *Square) Area() int {
	return s.side * s.side
}

func (r *Rectangle) Area() int {
	return r.length*r.breadth
}

func (h *Hybrid) Area() int {
	return h.s.Area() + h.r.Area()
}
