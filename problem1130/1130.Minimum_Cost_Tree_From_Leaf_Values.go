package problems

import "fmt"

// MctFromLeafValues problem 1130
func MctFromLeafValues(arr []int) int {
	return mctFromLeafValues(arr)
}

// Stack type
type Stack []int

func mctFromLeafValues(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	// init
	stack := make(Stack, 0)
	stackP := &stack
	top := stackP.push(arr[0])
	result := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] <= top {
			top = stackP.push(arr[i])
		} else {
			for arr[i] > top {
				if len(stack) > 1 {
					a := stackP.pop()
					b := stackP.pop()
					if b < arr[i] {
						result += a * b
						top = stackP.push(b)
					} else {
						result += a * arr[i]
						top = stackP.push(b)
						top = stackP.push(arr[i])
					}
				} else {
					a := stackP.pop()
					result += a * arr[i]
					top = stackP.push(arr[i])
				}
			}
		}
	}

	for len(stack) > 1 {
		a := stackP.pop()
		b := stackP.pop()
		result += a * b
		stackP.push(b)
	}

	return result
}

func (s *Stack) push(v int) int {
	fmt.Printf("%p, %p\n", s, *s)
	*s = append(*s, v)

	return v
}

func (s *Stack) pop() int {
	if len(*s) == 0 {
		return 0
	}
	len := len(*s)
	fmt.Println("len of stack", len)
	v := (*s)[len-1]
	*s = (*s)[:len-1]

	return v
}
