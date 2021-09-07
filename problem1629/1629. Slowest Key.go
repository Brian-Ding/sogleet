package problem1629

import "fmt"

func Judge() {
	releaseTimes := []int{9, 29, 49, 50}
	keysPressed := "cbcd"
	result := slowestKey(releaseTimes, keysPressed)

	fmt.Println(result)
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	time := releaseTimes[0]
	key := keysPressed[0]

	for i := 1; i < len(keysPressed); i++ {
		tempTime := releaseTimes[i] - releaseTimes[i-1]
		tempKey := keysPressed[i]

		if tempTime > time {
			time = tempTime
			key = tempKey
		} else if tempTime == time && tempKey > key {
			key = tempKey
		}
	}

	return key
}
