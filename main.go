package main

import (
	"fmt"
	"os"
	"regexp"
)

//NOTE -
// "fmt" is a GO library that is used to format basic strings, values, inputs, and outputs. [fmt package:] https://pkg.go.dev/fmt
// "regexp" is a GO package that implements regular expression search. [regexp package:] https://pkg.go.dev/regexp
// "os" is a package that provides a platform-independent interface to operating system functionality. [os package:] https://pkg.go.dev/os

//NOTE -
/*
| The main function in where the base code is typically stored
| We can also define a subset of "child" functions and execute them in the main function
*/

// NOTE -
// "var" means that its scope is global and can be accessed from elsewhere
// "const" just means that it is a constant

var restart string

// It starts a very simple routine that prompts inputs to solve basic arithmetic operations
func main() {
	// NOTE: we define the base calc function as a main child function: this is because we want to avoid to self recall main in case of errors
	startCalculation()
}

// This is the main routine that start a cascading sequence of basic input trigger.
// While the restart choice is set to "y", it will recursively recall self to start another operation
// Exits if the user's choice is different from "y" (yes)
func startCalculation() {
	// NOTE: selecting the operator and outputting the selction
	operator, operatorOutput := processOperator()
	fmt.Println(operatorOutput)

	// NOTE: selecting the two numbers and outputting the selection
	numbers, numberOutput := processNumbers()
	fmt.Println(numberOutput)

	// NOTE: return the result based on given operator
	result := processResult(operator, numbers[:])
	fmt.Println(result)

	// NOTE: asking if we have to start another operation
	fmt.Println("Would you like to perform another operation? [y/N]")
	fmt.Scan(&restart)

	if restart != "y" {
		fmt.Println("Application closed. Thank you for using this dumb calculator!")
		// NOTE: exit code 0 means that application has been closed without errors, 1 means closed with errors
		os.Exit(0)
	}
	startCalculation()

}

// It takes two numbers given by user
// The first param is a string representing the operator selected; the second param, is an array of float64 values
// It returns a formatted string as a single printed output of the entire arithmetic operation with its result
func processResult(operator string, numbers []float64) (output string) {
	fmt.Println(operator)
	switch operator {
	case "+":
		result := numbers[0] + numbers[1]
		output = fmt.Sprintf("%g %s %g = %g", numbers[0], operator, numbers[1], result)
		break
	case "-":
		result := numbers[0] + numbers[1]
		output = fmt.Sprintf("%g %s %g = %g", numbers[0], operator, numbers[1], result)
		break
	case "/":
		result := numbers[0] / numbers[1]
		output = fmt.Sprintf("%g %s %g = %g", numbers[0], operator, numbers[1], result)
		break
	case "*":
		result := numbers[0] * numbers[1]
		output = fmt.Sprintf("%g %s %g = %g", numbers[0], operator, numbers[1], result)
		break
	}
	return

}

// NOTE
//If a function has a named return (typed) value, we can use a naked return to let that value to be used later on
//We can also define multireturns simply by adding a second return value in named returns defined in function

// It checks the user's provided operator value: it calls the isAValidOperator() function to check if provided input is valid
// recall itself otherwise
func processOperator() (_operator, output string) {
	// this is the longhand definition of a variable
	var operator string
	// this is the shorthand definition of a variable
	// operator := 'some_value'
	fmt.Println("Enter operator: +, -, *, /")
	fmt.Scan(&operator)

	// we chack for the string operator provided
	if !isValidOperator(operator) {
		fmt.Println("The operator selected is invalid, please retry")
		return processOperator()
	}
	// %q is a verb that format the passed value within double quotes
	output = fmt.Sprintf("The selected operator is %q\n", operator)
	_operator = operator
	return _operator, output
}

// Ask for two float64 input values. When all values are provided, it also checks for its validity
// If they are invalid (ex. not float64 values), it recursively recall self to start the input process again
// Returns an array of selected input values and a string representing the log value
func processNumbers() (selectedNumbers [2]float64, output string) {
	var num1, num2 float64
	fmt.Println("Enter two numbers:")
	n, err := fmt.Scan(&num1, &num2)

	// NOTE: one of the return values of Scan is the number of items scanned: we know for sure it must be 2, so we can check against n value besides of err
	if err != nil && n < 2 {
		fmt.Println("One or more input is invalid, please retry")
		return processNumbers()
	}
	// NOTE:
	// verb %f is used for decimalles with just floating point and no exponent
	// verb %.2f specify the float type with a fixed number of 2 decimals
	// here we are printing two selected numbers with %g verb: we have already typed them as float but we don't know if they really are
	output = fmt.Sprintf("Selected numbers are %g and %g", num1, num2)
	selectedNumbers = [2]float64{num1, num2}
	return
}

// Check if given operator is valid, based on a regular expression that checks for four arithmetic's operations signs
// Return a bool output based on this check
func isValidOperator(operator string) (valid bool) {
	valid, err := regexp.MatchString(`^[\+,\-.\*,\/]$`, operator)
	if err != nil {
		errorMessage := fmt.Errorf(err.Error())
		fmt.Println(errorMessage)
	}
	return
}
