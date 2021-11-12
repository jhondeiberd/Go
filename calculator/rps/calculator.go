package rps

const (
	SUM            = 0
	REST           = 1
	MULTIPLICATION = 2
	DIVISION       = 3
	CALCULATOR     = 4
	CLOSE          = 5
)

type Operation struct {
	ResultCalc float64 `json:"resultCalc"`
	Test       string  `json:"test"`
	Number1    float64 `json:"number1"`
	Number2    float64 `json:"number2"`
}

//num2, _ := document.getElementById(number2).value

func Calculation(num1 float64, num2 float64, operator int) Operation {
	//num2 := strconv.ParseFloat(nm2, 64)
	resultCalc := 0.0
	var result Operation

	switch operator {
	case SUM:
		resultCalc = num1 + num2
		break
	case REST:
		resultCalc = num1 - num2
		break
	case MULTIPLICATION:
		resultCalc = num1 * num2
		break
	case DIVISION:
		resultCalc = num1 / num2
		break
	default:
	}
	result.ResultCalc = resultCalc
	//result.Nm1 = num1
	//result.Nm2 = num2
	return result
}
