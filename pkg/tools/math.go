package t

func CeilDivide(a, b int) int {
	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		// 同号情况
		return (a + b - 1) / b
	} else {
		// 异号情况
		return (a + b + 1) / b
	}
}

func CeilDivide64(a, b int64) int64 {
	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		// 同号情况
		return (a + b - 1) / b
	} else {
		// 异号情况
		return (a + b + 1) / b
	}
}

func CeilDivideUInt(a, b uint) uint {
	return (a + b - 1) / b
}
