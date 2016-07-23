package interfaces

type Square struct {
	side int
}

type Rectangle struct {
	length int
	breadth int
}

type Hybrid struct{
	s1 Shape
	s2 Shape
}

type Shape interface {
	Area() int
}

func (sq Square) Area() int{
	return sq.side*sq.side
}

func (rect Rectangle) Area() int{
	return rect.length*rect.breadth
}

func (hybrid Hybrid) Area() int{
	return hybrid.s1.Area() + hybrid.s2.Area()
}
