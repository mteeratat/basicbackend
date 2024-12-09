package calculator

func Calculator(mode string, num1, num2 float64) float64 {
	switch mode {
	case "plus":
		return num1 + num2
	case "minus":
		return num1 - num2
	case "multiply":
		return num1 * num2
	case "divide":
		return num1 / num2
	}
	return 0
}
