/*
Сколько лет вам было бы на Марсе?
Год на Земле состоит из 365 дней (високосные года не учитываем), а год на Марсе - из 687 дней.

Есть некая программа, которая принимает на вход ваш возраст в земных годах и выводит соответствующий возраст на Марсе.

Вам нужно написать для этой программы функцию mars_age(), которая будет принимать на вход целое число, возраст в земных 
годах, а возвращать возраст на Марсе. Возвращать функция должна целое число. Округлять ничего не надо. Просто отбрасываем дробную часть.

Sample Input:
1000
Sample Output:
531
*/

//For this tasks, we need to write only function, package "fmt" and func main () already declared.

func mars_age (earth int) int {
    mars := earth*365/687
    return mars
}
