/*
На вход подаются три целых числа. Необходимо сосчитать и вывести их сумму и произведение на разных строках.

Sample Input:
2 3 4

Sample Output:
9
24
*/

package main

import "fmt"

func main() {
// здесь вам нужно написать свой код
    var a, b, c int
    fmt.Scanln (&a, &b, &c)
    fmt.Println (a+b+c)
    fmt.Println (a*b*c)
}
