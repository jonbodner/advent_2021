package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

/*
You quickly make a list of the horizontal position of each crab (your puzzle input). Crab submarines have limited fuel, so you need to find a way to make all of their horizontal positions match while requiring them to spend as little fuel as possible.

For example, consider the following horizontal positions:

16,1,2,0,4,2,7,1,2,14
This means there's a crab with horizontal position 16, a crab with horizontal position 1, and so on.

Each change of 1 step in horizontal position of a single crab costs 1 fuel. You could choose any horizontal position to align them all on, but the one that costs the least fuel is horizontal position 2:

Move from 16 to 2: 14 fuel
Move from 1 to 2: 1 fuel
Move from 2 to 2: 0 fuel
Move from 0 to 2: 2 fuel
Move from 4 to 2: 2 fuel
Move from 2 to 2: 0 fuel
Move from 7 to 2: 5 fuel
Move from 1 to 2: 1 fuel
Move from 2 to 2: 0 fuel
Move from 14 to 2: 12 fuel
This costs a total of 37 fuel. This is the cheapest possible outcome; more expensive outcomes include aligning at position 1 (41 fuel), position 3 (39 fuel), or position 10 (71 fuel).

Determine the horizontal position that the crabs can align to using the least fuel possible. How much fuel must they spend to align to that position?
*/
func part1() {
	vals := getInitial()
	//vals = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	minTotal := math.MaxInt
	minPos := 0
	max := max(vals)
	for i := 0; i <= max; i++ {
		total := totalDistance(vals, i)
		if total < minTotal {
			minTotal = total
			minPos = i
		}
	}
	fmt.Println(minPos, minTotal)
}

/*
As it turns out, crab submarine engines don't burn fuel at a constant rate. Instead, each change of 1 step in horizontal position costs 1 more unit of fuel than the last: the first step costs 1, the second step costs 2, the third step costs 3, and so on.

As each crab moves, moving further becomes more expensive. This changes the best horizontal position to align them all on; in the example above, this becomes 5:

Move from 16 to 5: 66 fuel
Move from 1 to 5: 10 fuel
Move from 2 to 5: 6 fuel
Move from 0 to 5: 15 fuel
Move from 4 to 5: 1 fuel
Move from 2 to 5: 6 fuel
Move from 7 to 5: 3 fuel
Move from 1 to 5: 10 fuel
Move from 2 to 5: 6 fuel
Move from 14 to 5: 45 fuel
This costs a total of 168 fuel. This is the new cheapest possible outcome; the old alignment position (2) now costs 206 fuel instead.

Determine the horizontal position that the crabs can align to using the least fuel possible so they can make you an escape route! How much fuel must they spend to align to that position?
*/
func part2() {
	vals := getInitial()
	//vals = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	minTotal := math.MaxInt
	minPos := 0
	max := max(vals)
	for i := 0; i <= max; i++ {
		total := totalDistance2(vals, i)
		if total < minTotal {
			minTotal = total
			minPos = i
		}
	}
	fmt.Println(minPos, minTotal)
}

func max(vals []int) int {
	max := 0
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func totalDistance(vals []int, pos int) int {
	total := 0
	for _, v := range vals {
		total += int(math.Abs(float64(v - pos)))
	}
	return total
}

func totalDistance2(vals []int, pos int) int {
	total := 0
	for _, v := range vals {
		dist := int(math.Abs(float64(v - pos)))
		sum := float64(1+dist) * (float64(dist) / 2.0)
		total += int(sum)
	}
	return total
}

func getInitial() []int {
	contents, _ := os.ReadFile("./day7/input.txt")
	initial := strings.Split(string(contents), ",")
	//fmt.Println(initial)
	in := make([]int, len(initial), 1_000)
	for i := 0; i < len(initial); i++ {
		b, _ := strconv.Atoi(strings.TrimSpace(initial[i]))
		in[i] = b
	}
	return in
}
