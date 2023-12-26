package main

import (
	"fmt"
	"math"
	"strconv"
)

/// Week 3: Task 2

func main() {
	fmt.Println("start the program")
	getWorkWithUser()
	fmt.Println("end of program")
}

func getWorkWithUser() {
	for {
		if calculationPerforming() == "no" {
			break
		}
		calculation()
	}
}
func calculation() {
	a, b := getUserNumbers()
	operation := getUserOperation()
	res, err := calculate(a, b, operation)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("result is: %f \n", res)
	}
}

func calculationPerforming() string {
	for {
		fmt.Print("Do you want to perform calculation? (yes/no): ")
		var anotherCalculation string
		fmt.Scan(&anotherCalculation)

		if anotherCalculation == "yes" || anotherCalculation == "no" {
			return anotherCalculation
		}
		fmt.Printf("You enter the %s -it is not allowed", anotherCalculation)
	}

}

func calculate(a, b int, operation string) (float64, error) {
	var res float64
	var err error
	var r int
	switch operation {
	case "/":
		res, err = divide(a, b)
	case "*":
		r, err = multiply(a, b)
		res = float64(r)
	case "-":
		r, err = subtract(a, b)
		res = float64(r)
	case "+":
		r, err = sum(a, b)
		res = float64(r)
	case "^":
		r, err = exponentiation(a, b)
		res = float64(r)
	case "%":
		r, err = modulus(a, b)
		res = float64(r)
	default:
		res, err = 0, fmt.Errorf("operation not allowed")
	}
	return res, err
}

func validateNonZero(b int) error {
	if b == 0 {
		return fmt.Errorf("devision: devision by zero")
	}
	return nil
}

func getUserNumbers() (a int, b int) {

	for i := 0; i < 2; i++ {
		fmt.Printf("Type the %d number: ", i+1)
		var s string
		fmt.Scan(&s)
		a1, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			i--
		} else if i == 0 {
			a = a1
		} else {
			b = a1
		}
	}

	return
}

func getUserOperation() (operation string) {

	for {
		fmt.Println("Allowed operations: / * - + ^ %")
		fmt.Print("Type the operation:\n / for divide, * for multiply, - for subtract\n + for sum, ^ for exponentiation, '%' for modulus ")
		fmt.Scan(&operation)
		if !(operation == "/" || operation == "^" || operation == "%" || operation == "*" || operation == "-" || operation == "+") {
			fmt.Println("Operation not allowed")
		} else {
			return
		}
	}
}

func sum(a, b int) (int, error) {
	return a + b, nil
}

func multiply(a, b int) (int, error) {
	return a * b, nil
}

func subtract(a, b int) (int, error) {
	return a - b, nil
}

func divide(a, b int) (float64, error) {
	if err := validateNonZero(b); err != nil {
		return 0, err
	}
	return float64(a) / float64(b), nil
}

func exponentiation(a, b int) (int, error) {
	return int(math.Pow(float64(a), float64(b))), nil
}

func modulus(a, b int) (int, error) {
	var err error
	if err = validateNonZero(b); err != nil {
		return 0, err
	}

	return a % b, nil
}
