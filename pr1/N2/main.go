package main

import (
	"fmt"
	"regexp"
)

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

// функция для ввода только чисел
func inputInt(prompt string) int {
	var value int
	for {
		fmt.Println(prompt)
		_, err := fmt.Scanf("%d\n", &value)
		if err != nil {
			fmt.Println("Ошибка: нужно ввести только число!")
			var tmp string
			fmt.Scanln(&tmp) // очищаем ввод
			continue
		}
		return value
	}
}

// функция для ввода только букв
func inputLetters(prompt string) string {
	var value string
	re := regexp.MustCompile(`^[A-Za-zА-Яа-яЁё]+$`)
	for {
		fmt.Println(prompt)
		fmt.Scanf("%s\n", &value)
		if !re.MatchString(value) {
			fmt.Println("Ошибка: можно вводить только буквы!")
			continue
		}
		return value
	}
}

func main() {
	const size = 512
	empls := [size]*Employee{}

	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d\n", &cmd)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)
			empl.Name = inputLetters("Имя:")
			empl.Age = inputInt("Возраст:")
			empl.Position = inputLetters("Позиция:")
			empl.Salary = inputInt("Зарплата:")

			added := false
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					added = true
					fmt.Println("Сотрудник добавлен под номером", i)
					break
				}
			}
			if !added {
				fmt.Println("Ошибка: достигнут лимит в", size, "сотрудников")
			}

		case 2:
			// Удаляем сотрудника
			idx := inputInt("Введите номер сотрудника для удаления:")
			if idx >= 0 && idx < size && empls[idx] != nil {
				empls[idx] = nil
				fmt.Println("Сотрудник удалён")
			} else {
				fmt.Println("Ошибка: сотрудник с таким номером не найден")
			}

		case 3:
			// Выводим список сотрудников
			fmt.Println("\nСписок сотрудников:")
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					fmt.Printf("[%d] %s, %d лет, %s, зарплата: %d\n",
						i, empls[i].Name, empls[i].Age, empls[i].Position, empls[i].Salary)
				}
			}

		case 4:
			// Выход
			fmt.Println("Выход из программы...")
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
