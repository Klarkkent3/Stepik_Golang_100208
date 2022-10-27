/*
Напишите функцию max() которая должна принимать на вход два целых числа и возвращать большее из них.

Sample Input:
10 4

Sample Output:
10
*/

//For this tasks, we need to write only function, package "fmt" and func main () already declared.

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
