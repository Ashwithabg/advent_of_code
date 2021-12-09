package models

type RectangularPrism struct {
	Length int
	Width  int
	Height int
}

func (rectangularPrism RectangularPrism) FindSurfaceArea() int {
	return 3*rectangularPrism.Length*rectangularPrism.Width +
		2*rectangularPrism.Width*rectangularPrism.Height +
		2*rectangularPrism.Height*rectangularPrism.Length
}

func (rectangularPrism RectangularPrism) FindTotalLengthOfRibbonRequired() int {
	return 2*rectangularPrism.Length +
		2*rectangularPrism.Width +
		rectangularPrism.Length*rectangularPrism.Width*rectangularPrism.Height
}
