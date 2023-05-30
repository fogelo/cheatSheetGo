package main

import "fmt"

func main() {
	// demoType()
	// demoStruct()
	// demoMethod()
	demoStructMethod()
}

/*

 */

//@ 1. Демонстрация объявления типов

type mile int // объвили свой тип mile, который по является типом int (можно использовать такой подход для того чтобы указвать рамерности)

func demoType() {
	var x mile = 10
	fmt.Println(x)
}

//@ 2. Демонстрация объявления структур

type person struct { //создание новой структуры
	name string
	age  int
}

func demoStruct() {
	//инициализация структуры c помощью инициализатора (набор значений в фигурных скобках в соответствующем структуре порядке)
	var anton person = person{"anton", 32}
	var mary person = person{name: "mary", age: 32} //либо можно указывать поля явным образом
	fmt.Println(anton)
	fmt.Println(mary)
	fmt.Println(anton.name) //обращение к полям структур

	// как и в случае с обычными переменными можно создавать указатели на структуры
	var antonPointer *person = &anton
	(*antonPointer).age = 33
	fmt.Println(anton)

}

/*
Структуры можно вкладывать друг в друга. Можно создавать указатели на типы.
*/

//@ 3. Демонстрация создания и использования методов

/*

Метод - это функция, которая связана с определенным типов. С помощью имени параметра мы можем обращаться к свойствам структуры.
В остальномы методы это такие же функции

func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возпращаемых_результатов) {
	тело_метода
}
*/

type library []string //тип для среза

// метод для перебора элементов среза с типом library

var lib library = library{"book1, book2, book3"}

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

func demoMethod() {
	lib.print() //в данном случае lib и будет значением, которое передается в функцию через параметр l
}

//@ 4. Демонастрация создания и использования методов для структур

type man struct {
	name string
	age  int
}

func (m man) print2() {
	fmt.Println("имя: ", m.name)
	fmt.Println("возраст: ", m.age)
}
func (m man) eat(meal string) {
	fmt.Println(m.name, "ест", meal)
}

func demoStructMethod() {
	var anton man = man{"anton", 32}
	anton.print2()
	anton.eat("мясо")
}

//@ 5. Демонстрация создания и использования методов указателей
/*

При использовании метода, объект структуры, для которой определен метод передается по значению. Но можно использовать и указатели.

*/
