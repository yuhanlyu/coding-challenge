// You are given two jugs with capacities x and y litres. There is an infinite
// amount of water supply available. You need to determine whether it is
// possible to measure exactly z litres using these two jugs.

func canMeasureWater(x int, y int, z int) bool {
    if x + y < z {
        return false
    }
    if x + y == z {
        return true
    }
    for y != 0 {
        x, y = y, x % y
    }
    return z % x == 0
}
