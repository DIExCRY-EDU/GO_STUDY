package main

import (
	"fmt"
)

// видео 1:34:40
// про указатели подробнее прочитать в книге "Грокаем алгоритмы"

func main() {
	// message := "Скоро я буду Backend программистом"
	// fmt.Println(message)
	// printMessage(&message)
	// fmt.Println(message)

	number := 5
	var p *int = &number

	fmt.Println(*p)
	fmt.Println(&number)

	*p = 10
	fmt.Println(number)
}

// func printMessage(message *string) {
// 	*message += " (из функции printMessage() )"
// }

// name := "Max"
// age := 52
// new_m := sayHello(name, age)
// mes, err := checkAge(age)
// if err != nil {
// 	log.Fatal(err)
// }
// printMessage(new_m)
// printMessage(mes)

// dayOfweek := "asdad"
// fmt.Println(calendar(dayOfweek))

// fmt.Println(finMin(1, 23, 23432, 0, -1, 1241))

// func() {
// 	fmt.Println("Анонимная функция")
// }()

// inc := increment()
// fmt.Println(inc())
// fmt.Println(inc())
// fmt.Println(inc())
// fmt.Println(inc())

// func increment() func() string {
// 	count := 0
// 	return func() string {
// 		count++
// 		return fmt.Sprintf("Вызов функции increment(), count равен %d", count)
// 	}
// }

// func finMin(nums ...int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	min := nums[0]

// 	for _, i := range nums {
// 		if i < min {
// 			min = i
// 		}
// 	}

// 	return min
// }

// func calendar(day string) (string, error) {
// 	switch day {
// 	case "пн":
// 		return "Сегодня понедельник хорошего дня!", nil
// 	case "вт":
// 		return "Сегодня вторник хорошего дня!", nil
// 	case "ср":
// 		return "Сегодня среда хорошего дня!", nil
// 	case "чт":
// 		return "Сегодня четверг хорошего дня!", nil
// 	case "пт":
// 		return "Сегодня пятница хорошего дня и наступающих выходных!", nil
// 	case "сб":
// 		return "Сегодня суббота хороших выходных!", nil
// 	case "вс":
// 		return "Сегодня воскресенье, завтра на работу!", nil
// 	default:
// 		return "Неправильный ввод!", errors.New("incorrect input")
// 	}
// }

// func printMessage(message string) {
// 	fmt.Println(message)
// }

// func sayHello(name string, age int) string {
// 	return fmt.Sprintf("Привет, %s! Тебе %d лет.", name, age)
// }

// func checkAge(age int) (string, error) {
// 	if age >= 18 && age < 45 {
// 		return "Входи!", nil
// 	} else if age >= 45 {
// 		return "Не староват ли?", nil
// 	}

// 	return "Слышь тебе сюда нельзя!", errors.New("you are too young for this")
// }
