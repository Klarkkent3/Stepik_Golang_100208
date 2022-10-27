/*
Допишите код, чтобы программа считывала два целых числа со входа и выводила их сумму

Sample Input:
2
3

Sample Output:
5
*/

package main

import "fmt"

func main() {
    var a int
    var b int
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    fmt.Println(a + b)
}
