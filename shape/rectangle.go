package shape

import (
	"fmt"
)

//rectangle requirements
//positive values of p,w,h,d
//area calculation
// height * width
// diagonals^2 = height^2 + width^2
// height = width
//A=ab= P^2-4D^2/8

// type shape interface {
// 	CalculateArea()
// 	CalculatePerimeter()
// }

type rectangle struct {
	width  float64
	height float64
	// diagonal float64
}

// type square struct {
// 	rectangle
// }

type square struct {
	side float64
}

// func (rect rectangle) Side() float64 {

// 	return rect.width
// }

// type perimeter struct {
// 	p float64
// }

func RectangleValidity(w, h float64) (*rectangle, string) {
	if w > 0.0 && h > 0.0 {
		obj :=
			rectangle{
				width:  w,
				height: h,
			}
		return &obj, "It is a valid rectangle"
	} else {
		return nil, "Rectangle cannot have have zero or negative dimensions"
	}
}

func (r rectangle) CalculateArea() float64 {
	return r.width * r.height
}

// func (r rectangle) Area2() float64 {
// 	w := r.width
// 	d := r.diagonal
// 	h := math.Sqrt((d * d) - (w * w))

// 	return w * h
// }

// func (r rectangle) Area3() float64 {
// 	h := r.height
// 	d := r.diagonal
// 	w := math.Sqrt((d * d) - (h * h))

// 	return w * h
// }

func (r rectangle) CalculatePerimeter() float64 {
	return 2.0 * (r.height + r.width)
}

//	type Perimeter struct{
//		perimeter Perimeter1(),
//	}
// func (r rectangle) Perimeter2() float64 {
// 	return (2.0 * (r.height + math.Sqrt((r.diagonal*r.diagonal)-(r.height*r.height))))
// }

// func (r rectangle) Perimeter3() float64 {
// 	return (2.0 * (r.width + math.Sqrt((r.diagonal*r.diagonal)-(r.width*r.width))))
// }

// func (r rectangle) AreaFromPerimeterAndDiagonal(perimeter float64) float64 {
// 	d := r.diagonal
// 	a := ((perimeter * perimeter) - ((4.0) * (d * d))) / 8.0
// 	return a

// }

// func Square() {
// 	square := &rectangle{
// 	rectangle,
// 	}
// }

// func Area(side float64) float64 {
// 	// var side float64
// 	s := square{
// 		rectangle: rectangle{
// 			width: side,
// 		},
// 	}
// 	return s.rectangle.width * s.rectangle.width
// }

func Area(dimensions ...float64) string {
	if len(dimensions) == 1 {
		side := dimensions[0]
		return fmt.Sprintf("Area of square with side %.2f: %.2f", side, side*side)
	} else if len(dimensions) == 2 {
		if dimensions[0] == dimensions[1] {
			return fmt.Sprintf("Area of square with dimensions %.2f x %.2f: %.2f", dimensions[0], dimensions[1], dimensions[0]*dimensions[1])

		} else {
			rect := rectangle{
				width:  dimensions[0],
				height: dimensions[1],
			}
			return fmt.Sprintf("Area of rectangle with dimensions %.2f x %.2f: %.2f", rect.width, rect.height, rect.CalculateArea())
		}
	} else {
		return "Invalid number of dimensions provided"
	}
}

// 		return "These are the dimensions of a rectangle" // Return 0 for a rectangle
// 	} else {
// 		fmt.Println("Invalid number of dimensions provided")
// 		return 0
// 	}
// }

// func SquareValidity(s float64) (*square, string) {
// 	if s > 0.0 {
// 		obj :=
// 			square{
// 				side: s,
// 			}
// 		return &obj, "It is a valid square"
// 	} else {
// 		return nil, "Square cannot have have zero or negative dimensions"
// 	}
// }

// func (s square) CalculateArea() float64 {
// 	return s.side * s.side
// }

// func (s square) CalculatePerimeter() float64 {
// 	return 4.0 * s.side
// }
