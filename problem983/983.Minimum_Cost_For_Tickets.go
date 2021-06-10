package problems

/*
In a country popular for train travel, you have planned some train travelling one year in advance.
The days of the year that you will travel is given as an array days.  Each day is an integer from 1 to 365.

Train tickets are sold in 3 different ways:

a 1-day pass is sold for costs[0] dollars;
a 7-day pass is sold for costs[1] dollars;
a 30-day pass is sold for costs[2] dollars.
The passes allow that many days of consecutive travel.
For example, if we get a 7-day pass on day 2, then we can travel for 7 days: day 2, 3, 4, 5, 6, 7, and 8.

Return the minimum number of dollars you need to travel every day in the given list of days.
*/

// MincostTickets problem 983
func MincostTickets(days []int, costs []int) int {
	return mincostTickets(days, costs)
}

func mincostTickets(days []int, costs []int) int {
	row := len(days)
	array := make([]int, 0, row)
	for i := 0; i < row; i++ {
		array = append(array, 0)
	}
	array[row-1] = min(costs[0], costs[1], costs[2])

	for i := row - 2; i >= 0; i-- {
		next7DayIndex := -1
		next30DayIndex := -1

		// buy 1-day pass on day i
		temp := costs[0] + array[i+1]

		// buy 7-day pass on day i
		for j := i + 1; j < row; j++ {
			if days[i]+7 <= days[j] { // first day that 7-day pass does not work
				next7DayIndex = j
				break
			}
		}
		if next7DayIndex == -1 { // 7-day pass works until the end
			temp = min(temp, costs[1])
		} else {
			temp = min(temp, costs[1]+array[next7DayIndex])
		}

		// buy 30-day pass on day i
		for j := i + 1; j < row; j++ { // first day that 30-day pass does not work
			if days[i]+30 <= days[j] {
				next30DayIndex = j
				break
			}
		}
		if next30DayIndex == -1 { // 30-day pass works until the end
			temp = min(temp, costs[2])
		} else {
			temp = min(temp, costs[2]+array[next30DayIndex])
		}

		array[i] = temp
	}

	return array[0]
}
