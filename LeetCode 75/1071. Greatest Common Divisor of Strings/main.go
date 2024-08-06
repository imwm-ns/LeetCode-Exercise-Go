package main

import "fmt"

func main() {
	result := gcdOfStrings("ABCABC", "ABC")
	fmt.Println(result)
}

func gcdOfStrings(str1 string, str2 string) string {
	if (str1 + str2) == (str2 + str1) {
		return str1[0:gcd(len(str1), len(str2))]
	}
	return ""
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

/*
	Initialize
		The question asks, "What's the greatest common divisor of two strings?".
	So, we create a gcd function that requests two arguments, a and b, of int type,
	for solving the problem's greatest common divisor.

	then

	Input:
	a = "ABCABC", b = "ABC"
	a = 6, b = 3.

	round 1:
	Check b==0 (3==0). false =>   return gcd(3, 6% 3) = gcd(3, 0)

	round 2:
	Check b==0 (0==0). true => return a that is 3

	So, the result of this question is the character of str1 between 0's index and 3's index, which is "ABC.".
*/
