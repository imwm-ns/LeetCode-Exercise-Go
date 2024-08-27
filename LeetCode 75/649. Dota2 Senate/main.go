package main

import "fmt"

func main() {
	res := predictPartyVictory("RDR")
	fmt.Println(res)
}

func predictPartyVictory(senate string) string {
	queueRadiant, queueDire := make([]int, 0), make([]int, 0)

	for i, char := range senate {
		if char == 'R' {
			queueRadiant = append(queueRadiant, i)
		} else {
			queueDire = append(queueDire, i)
		}
	}

	n := len(senate)

	for len(queueRadiant) != 0 && len(queueDire) != 0 {
		if queueRadiant[0] < queueDire[0] {
			queueRadiant = append(queueRadiant, queueRadiant[0]+n)
		} else {
			queueDire = append(queueDire, queueDire[0]+n)
		}

		queueRadiant = queueRadiant[1:]
		queueDire = queueDire[1:]
	}

	if len(queueRadiant) != 0 {
		return "Radiant"
	}

	return "Dire"
}
