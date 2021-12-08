package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	part1()
	part2()
}

/*
Furthermore, you reason, a new lanternfish would surely need slightly longer before it's capable of producing more lanternfish: two more days for its first cycle.

So, suppose you have a lanternfish with an internal timer value of 3:

After one day, its internal timer would become 2.
After another day, its internal timer would become 1.
After another day, its internal timer would become 0.
After another day, its internal timer would reset to 6, and it would create a new lanternfish with an internal timer of 8.
After another day, the first lanternfish would have an internal timer of 5, and the second lanternfish would have an internal timer of 7.
A lanternfish that creates a new fish resets its timer to 6, not 7 (because 0 is included as a valid timer value). The new lanternfish starts with an internal timer of 8 and does not start counting down until the next day.

Realizing what you're trying to do, the submarine automatically produces a list of the ages of several hundred nearby lanternfish (your puzzle input). For example, suppose you were given the following list:

3,4,3,1,2
This list means that the first fish has an internal timer of 3, the second fish has an internal timer of 4, and so on until the fifth fish, which has an internal timer of 2. Simulating these fish over several days would proceed as follows:

Initial state: 3,4,3,1,2
After  1 day:  2,3,2,0,1
After  2 days: 1,2,1,6,0,8
After  3 days: 0,1,0,5,6,7,8
After  4 days: 6,0,6,4,5,6,7,8,8
After  5 days: 5,6,5,3,4,5,6,7,7,8
After  6 days: 4,5,4,2,3,4,5,6,6,7
After  7 days: 3,4,3,1,2,3,4,5,5,6
After  8 days: 2,3,2,0,1,2,3,4,4,5
After  9 days: 1,2,1,6,0,1,2,3,3,4,8
After 10 days: 0,1,0,5,6,0,1,2,2,3,7,8
After 11 days: 6,0,6,4,5,6,0,1,1,2,6,7,8,8,8
After 12 days: 5,6,5,3,4,5,6,0,0,1,5,6,7,7,7,8,8
After 13 days: 4,5,4,2,3,4,5,6,6,0,4,5,6,6,6,7,7,8,8
After 14 days: 3,4,3,1,2,3,4,5,5,6,3,4,5,5,5,6,6,7,7,8
After 15 days: 2,3,2,0,1,2,3,4,4,5,2,3,4,4,4,5,5,6,6,7
After 16 days: 1,2,1,6,0,1,2,3,3,4,1,2,3,3,3,4,4,5,5,6,8
After 17 days: 0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8
After 18 days: 6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8
Each day, a 0 becomes a 6 and adds a new 8 to the end of the list, while each other number decreases by 1 if it was present at the start of the day.

In this example, after 18 days, there are a total of 26 fish. After 80 days, there would be a total of 5934.

Find a way to simulate lanternfish. How many lanternfish would there be after 80 days?
*/
func part1() {
	in := getInitial()
	//fmt.Println(in)
	//in = []byte{3, 4, 3, 1, 2}
	for i := 0; i < 80; i++ {
		fmt.Println("day", i, ":", len(in))
		temp := make([]byte, 0, len(in))
		for _, v := range in {
			switch v {
			case 0:
				temp = append(temp, 8, 6)
			default:
				temp = append(temp, v-1)
			}
		}
		in = temp
	}
	fmt.Println(len(in))
}

func part2() {
	in := getInitial()
	//in = []byte{3, 4, 3, 1, 2}
	lookup := make([]int, 7)
	var wg sync.WaitGroup
	wg.Add(7)
	for i := 0; i <= 6; i++ {
		go func(i int) {
			start := time.Now()
			curSum := sumIt(i)
			lookup[i] = curSum
			fmt.Println(i, curSum, time.Now().Sub(start))
			wg.Done()
		}(i)
	}
	wg.Wait()
	total := len(in)
	for _, v := range in {
		total += lookup[v]
	}
	fmt.Println(total)
}

func sumIt(pos int) int {
	//fmt.Println("in sumIt starting at ", pos)
	made := int(math.Ceil((256 - float64(pos)) / 7))
	if made < 0 {
		return 0
	}
	//fmt.Println(made)
	total := made
	for i := 0; i <= made; i++ {
		p := pos + 9 + 7*i
		if p < 256 {
			total += sumIt(p)
		}
	}
	return total
}

func getInitial() []byte {
	contents, _ := os.ReadFile("./day6/input.txt")
	initial := strings.Split(string(contents), ",")
	//fmt.Println(initial)
	in := make([]byte, len(initial), 1_000)
	for i := 0; i < len(initial); i++ {
		b, _ := strconv.Atoi(strings.TrimSpace(initial[i]))
		in[i] = byte(b)
	}
	return in
}
