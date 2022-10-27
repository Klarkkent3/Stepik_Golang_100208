/*
На вход подается целое число, сумма денег, которые у вас есть. Ваша задача - вывести марку телефона, которую вы можете себе позволить купить.

Если сумма больше 1000 - вывести Apple
Если сумма от 500 до 1000 (включительно) - вывести Samsung
Если сумма меньше 500 - вывести Nokia с фонариком

Sample Input:
1100

Sample Output:
Apple
*/

package main

import "fmt"

func main() {
// здесь вам нужно написать свой код
    var sum int
    fmt.Scanln (&sum)
    switch {
        case sum > 1000:
        fmt.Println("Apple")
        case sum >= 500 && sum <= 1000:
        fmt.Println("Samsung")
        case sum < 500:
        fmt.Println("Nokia с фонариком")
    }   
}
