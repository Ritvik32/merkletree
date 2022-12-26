package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	input := []string{"A", "B", "C", "D"}
	var node []string

	if len(input)%2 != 0 {
		node = append(node, input[len(input)-1])
	}
	node = input
	for len(node) > 1 {
		var temp []string
		for j := 0; j < len(node); j += 2 {
			node1 := input[j]
			var node2 string
			if j+1 < len(input) {
				node2 = input[j+1]

			} else {
				temp = append(temp, node[j])
				break
			}
			//var a=[]byte(node1)
			//var b=[]byte(node2)
			var hashvalue []byte
			// var temp3 []byte
			temp1 := sha256.Sum256([]byte(node1))
			temp2 := sha256.Sum256([]byte(node2))
			//temp3 = append(temp3, temp1[:]...)
			//temp3 = append(temp3, temp2[:]...)
			hashvalue = append(hashvalue, temp1[:]...)
			hashvalue = append(hashvalue, temp2[:]...)
			temp3 := sha256.Sum256(hashvalue)

			mystring := string(temp3[:])
			temp = append(temp, mystring)

		}
		node = temp
	}

	fmt.Printf("%x", node)
}
