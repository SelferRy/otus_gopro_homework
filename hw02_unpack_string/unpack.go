package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func main() {
	input := ""
	expected := ""
	res, err := Unpack(input)
	if err != nil {
		fmt.Println("Failed")
	}
	fmt.Println(res)
	fmt.Println(expected)
}

func Unpack(data string) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	s := []rune(data)
	if firstIsDigit(s[0]) {
		return "", ErrInvalidString
	}
	var res strings.Builder
	count := len(s)
	for i := 0; i < count; i++ {
		switch {
		case lastElem(i, count):
			res.WriteString(string(s[i]))
		case matchPattern(s[i], s[i+1]):
			if isNumber(s, i, count) {
				return "", ErrInvalidString
			}
			times := takeTimes(s[i+1])
			res.WriteString(strings.Repeat(string(s[i]), times))
			i++
		default:
			res.WriteString(string(s[i]))
		}
	}
	return res.String(), nil
}

func firstIsDigit(first rune) bool {
	_, err := strconv.Atoi(string(first))
	return err == nil
}

func lastElem(i int, count int) bool {
	return i == count-1
}

func matchPattern(curr rune, next rune) bool {
	_, currIsDigit := strconv.Atoi(string(curr))
	_, nextIsDigit := strconv.Atoi(string(next))
	return !isDigit(currIsDigit) && isDigit(nextIsDigit)
}

func isNumber(s []rune, i int, count int) bool {
	if i+2 < count {
		_, errSecond := strconv.Atoi(string(s[i+1]))
		_, errThird := strconv.Atoi(string(s[i+2]))
		return isDigit(errSecond) && isDigit(errThird)
	}
	return false // length is too short
}

func isDigit(err error) bool {
	return err == nil
}

func takeTimes(num rune) int {
	n, _ := strconv.Atoi(string(num))
	return n
}
