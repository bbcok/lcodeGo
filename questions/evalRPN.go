package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		val, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, val)
		} else {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			switch tokens[i] {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}

		}

	}
	return stack[0]
}
