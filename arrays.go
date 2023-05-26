// § Массивы

package main

import "fmt"

func main() {
	var numbers1 [5]int
	numbers1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers1)

	numbers2 := [3]int{1, 2, 3}
	numbers2 = [3]int{4, 5, 6} //переприсовение
	fmt.Println(numbers2)

	var numbers3 = [...]int{1, 2, 3} // если указано троеточие, то размер массива определяется исходя из длины присовенного массива
	numbers3 = [3]int{4, 5, 6}
	fmt.Println(numbers3)

	// обращение к элементам массива через индексы
	fmt.Println(numbers3[0])
}

/*
?Мои замечания
1) Нужно задавать тип и переменной при объявлении и значению при присвоении, либо только при присвоениии
2) Если в типе указан элемент массива, но не был присвоен, то по умолчанию ему присвоится значение 0. Поэтому можно при присвоении указать меньше значений,
но больше нельзя

*/