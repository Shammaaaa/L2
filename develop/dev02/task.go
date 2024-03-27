package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi")
}

func unpack(line string) (string, error) {
	unpacked := []string{}

	escape := false

	s := strings.Split(line, "")

	for i, r := range s {
		if r == "\\" {
			if escape {
				unpacked = append(unpacked, r)
				escape = false
				continue
			}
			escape = true
			continue
		}

		isDigit, err := regexp.MatchString("[1-9]", r)
		if err != nil {
			return "", err
		}

		if i == 0 && isDigit {
			return "", fmt.Errorf("error: invalid line")
		}

		if isDigit {
			if !escape {
				n, err := strconv.Atoi(r)
				if err != nil {
					return "", err
				}

				unpacked = append(unpacked, strings.Repeat(s[i-1], n-1))
				continue
			}

			unpacked = append(unpacked, r)
			escape = false
		}

		isLetter, err := regexp.MatchString("[a-zA-Z]", r)
		if err != nil {
			return "", err
		}

		if isLetter {
			unpacked = append(unpacked, r)
			continue
		}
	}

	return strings.Join(unpacked, ""), nil
}
