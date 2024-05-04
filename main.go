package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	name := "Max"
	age := 52
	new_m := sayHello(name, age)
	mes, err := checkAge(age)
	if err != nil {
		log.Fatal(err)
	}
	printMessage(new_m)
	printMessage(mes)

}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет.", name, age)
}

func checkAge(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Входи!", nil
	} else if age >= 45 {
		return "Не староват ли?", nil
	}

	return "Слышь тебе сюда нельзя!", errors.New("you are too young for this")
}

// видео 1:19:00
