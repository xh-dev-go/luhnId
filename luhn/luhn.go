package luhn

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func toNumber(i int32) int32 {
	return i - 48
}

func sumAll(str string) int32 {
	var sum int32 = 0
	for _, vv := range str {
		sum += toNumber(vv)
	}
	return sum
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func getCheckDigit(str string) int32 {
	code := reverse([]rune(str))
	var sum int32 = 0
	for index, value := range code {
		v := toNumber(value)
		if index%2 == 0 {
			sum += sumAll(fmt.Sprintf("%02d", v*2))
		} else {
			sum += v
		}
	}

	mod := sum % 10
	if mod == 0 {
		return 0
	} else {
		return 10 - mod
	}
}

func Validate(str string) bool {
	checkDigit := toNumber([]rune(str)[len(str)-1])
	return getCheckDigit(str[0:len(str)-1]) == checkDigit
}

var err = errors.New("digit should not be zero or less")

func Gen(startWith string, digit int) (string, error) {
	var s = ""
	var d = 0
	if len(startWith) != 0 {
		d = digit - 1
	}

	if d <= 0 {
		return s, err
	}
	rand.Seed(time.Now().UTC().UnixNano())
	s = startWith
	for i := 0; i < d-1; i++ {
		s += fmt.Sprintf("%d", rand.Intn(10))
	}
	s += fmt.Sprintf("%d", getCheckDigit(s))
	return s, nil
}
