package main

import "fmt"

func main() {
	arr := []byte{'g', 'e', 'e', 'k', 's', 'f', 'o', 'r', 'g', 'e', 'e', 'k', 's'}
	fmt.Println(string(sortASCII(arr)))
}

func sortASCII(arr []byte) []byte {
	var frequency [128]int
	for i := 0; i < len(arr); i++ {
		frequency[arr[i]]++
	}
	j := 0
	for i := 0; i < 128; i++ {
		for ; frequency[i] > 0; frequency[i]-- {
			arr[j] = byte(i)
			j++
		}
	}
	return arr
}
