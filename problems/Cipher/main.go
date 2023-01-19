package main

import "strings"

func Atbash(s string) string {
	var shiftedStr string
  i:=0
	for _, r := range s {
		base := 0
		skipSpace := false
		switch {
		case r >= 'a' && r <= 'z':
			base = int('a')
		case r >= 'A' && r <= 'Z':
			base = int('A')
		case r >= '0' && r <= '9':
		default:
			skipSpace = true
		}
		if base != 0 {
			shift := 25 - (int(r) - base)
			r = rune(base + shift)
		}
		if !skipSpace {
			if i%5 == 0 && i > 0 {
				shiftedStr += " "
			}
			shiftedStr += string(r)
			i++
		}
	}
	return strings.ToLower(shiftedStr)
}

func main() {

}
