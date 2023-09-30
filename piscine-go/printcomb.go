package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for c := j + 1; c <= '9'; c++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(c)
				if i == '7' && j == '8' && c == '9' {
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
