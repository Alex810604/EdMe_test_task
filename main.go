package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) + 1
	attempt := 1 //  счетчик попыток
	num := 3     //  количество попыток
	fmt.Println("Угадай число от 1 до 10!")
	fmt.Printf("У вас %d попыток.", num)
	for {
		if attempt > num {
			fmt.Println("Ваши попытки закончились. Вы число не угадали.")
			break
		}
		var maybe int // возможное число, число попыток
		fmt.Printf("Попытка %d.\n", attempt)
		fmt.Print("Ваш вариант: ")
		fmt.Scan(&maybe)

		if maybe < 1 || maybe > 10 {
			fmt.Println("Ошибка! Введенное число не попадает в диапазон от 1 до 10!")
			continue
		}

		if maybe > n {
			fmt.Println("Загаданное число меньше!")
		} else if maybe < n {
			fmt.Println("Загаданное число больше!")
		} else {
			fmt.Printf("Вы угадали! Это %d.", maybe)
			break
		}
		attempt++
	}

}
