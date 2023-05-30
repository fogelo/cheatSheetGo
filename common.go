package main

import "fmt"

/*
?Конспект
go run main.go - компилирование и выполнение программы
go build main.go - компилирование программы (создание бинарного файла) (чтобы запустить ее потом введи ./main)

Типы данных:
		string
		int
		float32
		bool
		byte
лучше сразу привыкать в go работать с байтами (byte). Из байтов состоит все отстальное
*/

/*
?Вопросы:
1) Как использовать debug console?
2) Как запустить программу горячими клавишами?
3) Какие есть типы переменных в go?

*/

//@ Пример 2
func main() {
	text1 := "Hello"
	text2 := "World"

	// i := 10
	// так делать нельзя, поскольку язык с сильной типизацией и в нем соответственно нет приведения типов
	// text := text1 + i

	fmt.Println(text1, text2)
}
