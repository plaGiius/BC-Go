package shape

import (
	"testing"
)

func TestCheckingValidityOfRectangleProperties(t *testing.T) {
	t.Run("Validating the rectangle", func(t *testing.T) {

		ptr, _ := RectangleValidity(8.0, -1.0)
		// expected := "valid"

		if ptr == nil {
			t.Errorf("The rectangle can only have non-zero positive dimensions")
		}
	})

}

func TestRectangleAreaWhenWidthIs8AndHeightIs6(t *testing.T) {
	t.Run("Finding the area of rectangle", func(t *testing.T) {
		//setup

		w := 8.0
		h := 6.0

		rectangleTest := rectangle{
			width:  w,
			height: h,
		}

		area := rectangleTest.CalculateArea()
		expected := 48.0

		if area != expected {
			t.Errorf("expected '%f' but got '%f'", expected, area)
		}
	})
}

// func TestRectangleArea2(t *testing.T) {
// 	t.Run("Finding the area of rectangle when diagonal and width is given", func(t *testing.T) {
// 		w := 8.0
// 		d := 10.0

// 		rectangleTest := rectangle{
// 			width:    w,
// 			diagonal: d,
// 		}

// 		area := rectangleTest.Area2()
// 		expected := 48.0

// 		if area != expected {
// 			t.Errorf("expected '%f' but got '%f'", expected, area)
// 		}
// 	})
// }

// func TestRectangleArea3(t *testing.T) {
// 	t.Run("Finding the area of rectangle when diagonal and height is given", func(t *testing.T) {
// 		h := 6.0
// 		d := 10.0

// 		rectangleTest := rectangle{
// 			height:   h,
// 			diagonal: d,
// 		}

// 		area := rectangleTest.Area3()
// 		expected := 48.0

// 		if area != expected {
// 			t.Errorf("expected '%f' but got '%f'", expected, area)
// 		}
// 	})
// }

func TestRectanglePerimeterWhenWidthIs8AndHeightIs6(t *testing.T) {
	t.Run("Finding the perimeter of rectangle when width and height is given", func(t *testing.T) {
		w := 8.0
		h := 6.0

		rectangleTest := rectangle{
			height: h,
			width:  w,
		}

		perimeter := rectangleTest.CalculatePerimeter()
		expected := 28.0

		if perimeter != expected {
			t.Errorf("expected '%f' but got '%f'", expected, perimeter)
		}
	})
}

// func TestRectanglePerimeter2(t *testing.T) {
// 	t.Run("Finding the perimeter of rectangle when height and diagonal is given", func(t *testing.T) {
// 		h := 6.0
// 		d := 10.0

// 		rectangleTest := rectangle{
// 			height:   h,
// 			diagonal: d,
// 		}

// 		perimeter := rectangleTest.Perimeter2()
// 		expected := 28.0

// 		if perimeter != expected {
// 			t.Errorf("expected '%f' but got '%f'", expected, perimeter)
// 		}
// 	})
// }

// func TestRectanglePerimeter3(t *testing.T) {
// 	t.Run("Finding the perimeter of rectangle when height and diagonal is given", func(t *testing.T) {
// 		w := 8.0
// 		d := 10.0

// 		rectangleTest := rectangle{
// 			width:    w,
// 			diagonal: d,
// 		}

// 		perimeter := rectangleTest.Perimeter3()
// 		expected := 28.0

// 		if perimeter != expected {
// 			t.Errorf("expected '%f' but got '%f'", expected, perimeter)
// 		}
// 	})
// }

// func TestCalculateAreaFromPerimeterAndDiagonal(t *testing.T) {
// 	t.Run("Finding the area of rectangle when perimeter and diagonal is given", func(t *testing.T) {
// 		d := 10.0

// 		rect := rectangle{width: 0, height: 0, diagonal: d} // Set only diagonal

// 		// Calculate perimeter before calling the method
// 		perimeter := 28.0

// 		area := rect.AreaFromPerimeterAndDiagonal(perimeter) // Use perimeter as argument

// 		expected := 48.0

// 		if area != expected {
// 			t.Errorf("expected '%f' but got '%f'", expected, area)
// 		}
// 	})
// }

// func TestCheckingValidityOfSquareProperties(t *testing.T) {
// 	t.Run("Validating the square", func(t *testing.T) {

// 		ptr, _ := SquareValidity(-1.0)
// 		// expected := "valid"

// 		if ptr == nil {
// 			t.Errorf("The square can only have non-zero positive dimensions")
// 		}
// 	})

// }

// func TestSquareAreaWhenSideIs10(t *testing.T) {
// 	t.Run("Finding the area of square given a side", func(t *testing.T) {
// 		//setup

// 		s := 10.0

// 		squareTest := square{
// 			side: s,
// 		}

// 		area := squareTest.CalculateArea()
// 		expected := 100.0

// 		if area != expected {
// 			t.Errorf("expected '%f' but got '%f'", expected, area)
// 		}
// 	})
// }

// func TestSquarePerimeterWhenSideIs10(t *testing.T) {
// 	t.Run("Finding the perimeter of rectangle when width and height is given", func(t *testing.T) {
// 		s := 10.0

// 		squareTest := square{
// 			side: s,
// 		}

// 		perimeter := squareTest.CalculatePerimeter()
// 		expected := 40.0

// 		if perimeter != expected {
// 			t.Errorf("expected '%f' but got '%f'", expected, perimeter)
// 		}
// 	})
// }

// func TestCalculateArea(t *testing.T) {
// 	// Mocking the square struct
// 	type mockSquare struct {
// 		rectangle
// 	}

// 	mockSide := 5.0 // Setting a mock side length

// 	// Creating a mock square instance
// 	mockSquareInstance := mockSquare{
// 		rectangle: rectangle{
// 			width:  mockSide,
// 			height: mockSide,
// 		},
// 	}

// 	// Defining the expected area
// 	expectedArea := mockSide * mockSide

// 	// Calculating the area
// 	calculatedArea := Area()

// 	// Checking if the calculated area matches the expected area
// 	if calculatedArea != expectedArea {
// 		t.Errorf("Incorrect area calculation, got: %f, want: %f", calculatedArea, expectedArea)
// 	}
// }

// func TestCalculateArea(t *testing.T) {
// 	// Defining a side length
// 	s := square{
// 		rectangle: rectangle{
// 			width: 10.0,
// 		},
// 	}
// 	expectedArea := s.Side() * s.Side()

// 	// Expected area

// 	// Calculating the area
// 	calculatedArea := Area()

//		// Checking if the calculated area matches the expected area
//		if calculatedArea != expectedArea {
//			t.Errorf("Incorrect area calculation, got: %f, want: %f", calculatedArea, expectedArea)
//		}
//	}
//
// rectangle_test.go
// func TestCalculateArea(t *testing.T) {
// 	// Defining a side length
// 	Side := 5.0

// 	// Expected area
// 	expectedArea := Side * Side

// 	// Calculating the area
// 	calculatedArea := Area(Side)

// 	// Checking if the calculated area matches the expected area
// 	if calculatedArea != expectedArea {
// 		t.Errorf("Incorrect area calculation, got: %f, want: %f", calculatedArea, expectedArea)
// 	}
// }

// func TestArea(t *testing.T) {
// 	testSquare := func(side float64, expected string) {
// 		t.Helper()
// 		result := Area(side)
// 		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
// 			t.Errorf("For side %.2f, expected: %v, but got: %v", side, expected, result)
// 		}
// 	}

// 	testRectangle := func(width, height float64, expected string) {
// 		t.Helper()
// 		result := Area(width, height)
// 		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
// 			t.Errorf("For dimensions %.2f x %.2f, expected: %v, but got: %v", width, height, expected, result)
// 		}
// 	}

// 	// Test cases for squares
// 	testSquare(5.0, "Area of square with side 5.00: 25.00")
// 	testSquare(10.0, "Area of square with side 10.00: 100.00")

// 	// Test cases for rectangles
// 	testRectangle(5.0, 4.0, "Area of rectangle with dimensions 5.00 x 4.00: 20.00")
// 	testRectangle(8.0, 6.0, "Area of rectangle with dimensions 8.00 x 6.00: 48.00")

// 	// Test case for invalid input
// 	result := Area(5.0, 4.0, 3.0)
// 	expected := "Invalid number of dimensions provided"
// 	if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
// 		t.Errorf("For invalid input, expected: %v, but got: %v", expected, result)
// 	}
// }
