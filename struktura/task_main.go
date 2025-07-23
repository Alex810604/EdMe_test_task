package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Генерация случайного числа от 1 до 10
	n := rand.Intn(10) + 1 //Для Go ≥ 1.22 можно сразу
	attempt := 1           //  счетчик попыток
	num := 3               //  количество попыток
	fmt.Println("Угадай число от 1 до 10!")
	fmt.Printf("У вас %d попыток.", num)
	for {
		if attempt > num {
			fmt.Printf("Попытки закончились. Вы число не угадали. Число было %d.\n", n)
			break
		}
		var maybe int // возможное число
		fmt.Printf("Попытка %d.\n", attempt)
		fmt.Print("Ваш вариант: ")
		_, err := fmt.Scan(&maybe)
		if err != nil {
			fmt.Println("Ошибка!", err)
			continue
		}

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
