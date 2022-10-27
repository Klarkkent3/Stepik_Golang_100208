/*
Напишите функцию calc() которая должна принимать на вход одно целое число, а возвращать два значения - число умноженное на два и число возведенное в квадрат.

Sample Input:
4

Sample Output:
8 16
*/

//For this tasks, we need to write only function, package "fmt" and func main () already declared.

func calc (x int) (int,int) {
    return x*2, x*x
}
