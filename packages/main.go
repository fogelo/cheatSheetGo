package main

import (
	"fmt"
)

//§ пакеты, модули и их импорт
/*

Весь код в go организуется в пакеты..Пакеты представлюят удобную организацию и разделение кода на отдельные части или модули.
Модульность позволяет определять один раз пакет с нужной функциональностью и потом использовать его многократно в различных программах.

Есть 2 типа пакетов:
1. Исполняемые (executable) - должны иметь имя main и одноименную функцию main, которая является входной точкой в приложение
2. Библиотеки (reusable)


Мы можем импортировать как встроенные пакеты так и свои собственные. Один пакет может состоять из нескольских файлов.

Чтобы функция была видна в других пакетах ее имя должно начинаться с заглавной буквы.
*/

func main() {
	fmt.Println("hello wrold")
	hello()
}
