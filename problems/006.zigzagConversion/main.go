package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	result := string(s[0])
	up := true
	num := 0

	for curNum := 1; curNum <= numRows; {
		if up {
			num += (numRows - curNum) * 2
		} else {
			num += (curNum - 1) * 2
		}

		if num >= len(s) {
			up = true
			curNum++
			num = curNum - 1
			if num < numRows {
				result += string(s[num])
			}
		} else {
			if (up && numRows != curNum) || (!up && curNum != 1) {
				result += string(s[num])
			}
			up = !up
		}
	}

	return result
}
