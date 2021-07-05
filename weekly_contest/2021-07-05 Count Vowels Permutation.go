package weekly_contest

import "fmt"

func Judge() {
	var result int = countVowelPermutation(5)
	fmt.Println(result)
}

func countVowelPermutation(n int) int {
	records := make(map[rune]int)

	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 5
	}

	records['a'] = 1
	records['e'] = 1
	records['i'] = 1
	records['o'] = 1
	records['u'] = 1

	for i := 2; i <= n; i++ {
		a := records['a']
		e := records['e']
		i := records['i']
		o := records['o']
		u := records['u']

		for j := 0; j < a; j++ {
			permutate('a', records)
		}

		for j := 0; j < e; j++ {
			permutate('e', records)
		}

		for j := 0; j < i; j++ {
			permutate('i', records)
		}

		for j := 0; j < o; j++ {
			permutate('o', records)
		}

		for j := 0; j < u; j++ {
			permutate('u', records)
		}
	}

	sum := 0
	for _, v := range records {
		sum += v
	}

	return sum
}

func permutate(c rune, records map[rune]int) {
	if records[c] == 0 {
		return
	}

	records[c]--
	switch c {
	case 'a':
		records['e']++
	case 'e':
		records['a']++
		records['i']++
	case 'i':
		records['a']++
		records['e']++
		records['o']++
		records['u']++
	case 'o':
		records['i']++
		records['u']++
	case 'u':
		records['a']++
	}
}
