type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	return ""
}

func spiralOrder(matrix [][]int) []int {
	w := len(matrix[0])
	h := len(matrix)

	order := make([]int, 0, w*h)
	var minX, minY int

	maxX, maxY := w-1, h-1

	dir := Right
	for {
		switch dir {
		case Right:
			for x, y := minX, minY; x <= maxX; x++ {
				order = append(order, matrix[y][x])
			}
			minY++

			dir = Down
		case Down:
			for x, y := maxX, minY; y <= maxY; y++ {
				order = append(order, matrix[y][x])
			}
			maxX--
			dir = Left
		case Left:
			for x, y := maxX, maxY; x >= minX; x-- {
				order = append(order, matrix[y][x])
			}
			maxY--
			dir = Up
		case Up:
			for x, y := minX, maxY; y >= minY; y-- {
				order = append(order, matrix[y][x])
			}
			minX++
			dir = Right
		}

		// fmt.Printf("Next %s: minX=%d, maxX=%d minY=%d, maxY=%d\n", dir, minX, maxX, minY, maxY)
		// fmt.Println(order)
		// fmt.Println("expected=[1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]")
		// fmt.Println("len(order)", len(order), "w*h", w*h)

		if len(order) >= w*h {
			return order
		}

	}
	return order
}


