/*
Напишите функцию isEven(), которая принимает в качестве аргумента одно целое число и возвращает true если оно четное и false в ином случае.

Sample Input:
8

Sample Output:
true
*/

//For this tasks, we need to write only function, package "fmt" and func main () already declared.

func isEven(a int) bool {
	if a%2 != 0 {
		return false
	} else {
		return true
	}
}
