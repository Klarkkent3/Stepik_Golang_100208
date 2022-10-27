/*
Вам нужно написать функцию one_or_two(),  которая принимает два целых числа и строку. Строка может иметь значения one, two 
или любой другой текст. 

Возвращать из функции вам нужно два значения. Если строка равна one, нужно вернуть первое число и саму строку. Если строка 
равна two, нужно вернуть второе число и саму строку. Если строка другая - нужно вернуть 0 и саму строку.

Sample Input:

2 5 two
Sample Output:

5 two
*/

//For this tasks, we need to write only function, package "fmt" and func main () already declared.

// здесь вам нужно написать свой код
func one_or_two (a int, b int, c string) (int, string) {
    if c == "one" {
        return (a),("one")
    } else if c == "two" {
        return (b),("two")
    } else {
        return (0),(c)
    }
}
