type Direction int

const (
    Right Direction = iota
    Down
    Left
    Up
)

func generateMatrix(n int) [][]int {
    nums := make([][]int, n)
    for i:=0; i< n; i++ {
        nums[i] = make([]int, n)
    }


    var minX, minY int
    maxX, maxY := n-1, n-1
    dir := Right
    stepCount :=1
    for {
        switch dir {
            case Right:
            for x,y := minX, minY; x <= maxX; x++{
                nums[y][x] = stepCount
                stepCount++
            }
            minY++
            dir = Down
            case Down:
            for x,y := maxX, minY; y <= maxY; y++{
                nums[y][x] = stepCount
                stepCount++
            }
            maxX--
            dir = Left
            case Left:
            for x,y := maxX, maxY; x >= minX; x--{
                nums[y][x] = stepCount
                stepCount++
            }
            maxY--
            dir = Up
            case Up:
            for x,y := minX, maxY; y >= minY; y--{
                nums[y][x] = stepCount
                stepCount++
            }
            minX++
            dir = Right
        }

        if stepCount > n*n {
            return nums
        }
    }

    return nums
}