/*
Напишите функцию double_m(), которая должна принимать на вход два целых числа a и b и возвращать сумму квадратов чисел от
a до b (включительно).

Sample Input:
2 6
Sample Output:
90
*/

//For this tasks, we need to write only function, package "fmt" and func main () already declared.

func double_m (a, b int) int{
    sum := 0
    for  i := a; i<=b; i++ {
    sum += i*i
    }
    return sum
}

