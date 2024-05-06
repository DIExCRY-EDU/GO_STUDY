package main

import "fmt"

// видео 2:29:00
// про указатели подробнее прочитать в книге "Грокаем алгоритмы"
// effective go

func main() {
	matrix := make([][]int, 10)

	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		matrix[j] = make([]int, 10)
	// 		matrix[j][i] = 1
	// 	}
	// 	fmt.Println(matrix[i])
	// }

	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)
		matrix[i][i] = 1
		fmt.Println(matrix[i])
	}

}

// var messages [3]string = [3]string{"1", "2", "3"} // messages := [3]string{"1", "2", "3"}

// messages := make([]string, 100)
// messages = append(messages, "101")
// messages = append(messages, "7")
// messages = append(messages, "8")
// messages = append(messages, "9")
// messages = append(messages, "10")
// fmt.Println(len(messages))
// fmt.Println(cap(messages))

// func printMessages(messages []string) error {
// 	if len(messages) == 0 {
// 		return errors.New("empty arr of messages")
// 	}
// 	messages[1] = "5"
// 	fmt.Println(messages)
// 	return nil
// }

// message := "Скоро я буду Backend программистом"
// fmt.Println(message)
// printMessage(&message)
// fmt.Println(message)

// number := 5
// var p *int = &number

// fmt.Println(*p)
// fmt.Println(&number)

// *p = 10
// fmt.Println(number)

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
