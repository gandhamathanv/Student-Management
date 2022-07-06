package components

import (
	"fmt"
)
func Choice(placeholder string)int{
	var choice int
	fmt.Println(placeholder)
	fmt.Scanln(&choice)
	return choice
}
func Info(placeholder string)string{
	var data string
	fmt.Println(placeholder)
	fmt.Scanln(&data)
	return data
}