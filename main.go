package main

import (
	"fmt"
)

func main() {
	name := "Max"
	age := 16
	new_m := sayHello(name, age)
	mes, _ := checkAge(age)
	printMessage(new_m)
	printMessage(mes)

}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет.", name, age)
}

func checkAge(age int) (string, bool) {
	if age >= 18 {
		return "Входи!", true
	}

	return "Слышь тебе сюда нельзя!", false
}

// видео 1:10:00
