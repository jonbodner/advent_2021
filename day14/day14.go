package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	part1()
	part2()
}

/*
The submarine manual contains instructions for finding the optimal polymer formula; specifically, it offers a polymer template and a list of pair insertion rules (your puzzle input). You just need to work out what polymer would result after repeating the pair insertion process a few times.

For example:

NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
The first line is the polymer template - this is the starting point of the process.

The following section defines the pair insertion rules. A rule like AB -> C means that when elements A and B are immediately adjacent, element C should be inserted between them. These insertions all happen simultaneously.

So, starting with the polymer template NNCB, the first step simultaneously considers all three pairs:

The first pair (NN) matches the rule NN -> C, so element C is inserted between the first N and the second N.
The second pair (NC) matches the rule NC -> B, so element B is inserted between the N and the C.
The third pair (CB) matches the rule CB -> H, so element H is inserted between the C and the B.
Note that these pairs overlap: the second element of one pair is the first element of the next pair. Also, because all pairs are considered simultaneously, inserted elements are not considered to be part of a pair until the next step.

After the first step of this process, the polymer becomes NCNBCHB.

Here are the results of a few steps using the above rules:

Template:     NNCB
After step 1: NCNBCHB
After step 2: NBCCNBBBCBHCB
After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
This polymer grows quickly. After step 5, it has length 97; After step 10, it has length 3073. After step 10, B occurs 1749 times, C occurs 298 times, H occurs 161 times, and N occurs 865 times; taking the quantity of the most common element (B, 1749) and subtracting the quantity of the least common element (H, 161) produces 1749 - 161 = 1588.

Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result. What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?
*/
func part1() {
	data := buildData()
	start := time.Now()
	for i := 0; i < 10; i++ {
		newRow := make([]byte, len(data.curChain)*2-1)
		for i := 0; i < len(data.curChain)-1; i++ {
			newRow[i*2] = data.curChain[i]
			s := data.curChain[i : i+2]
			newRow[i*2+1] = byte(data.rules[s])
		}
		newRow[len(newRow)-1] = data.curChain[len(data.curChain)-1]
		data.curChain = string(newRow)
	}
	counts := calcCounts(data.curChain)
	fmt.Println(time.Since(start))
	minCount := math.MaxInt
	maxCount := 0
	for _, v := range counts {
		if v > maxCount {
			maxCount = v
		}
		if v < minCount {
			minCount = v
		}
	}
	fmt.Println(maxCount, minCount, maxCount-minCount)
}

func calcCounts(s string) map[rune]int {
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	return m
}

/*
The resulting polymer isn't nearly strong enough to reinforce the submarine. You'll need to run more steps of the pair
insertion process; a total of 40 steps should do it.

In the above example, the most common element is B (occurring 2192039569602 times) and the least common element is H
(occurring 3849876073 times); subtracting these produces 2188189693529.

Apply 40 steps of pair insertion to the polymer template and find the most and least common elements in the result.
What do you get if you take the quantity of the most common element and subtract the quantity of the least common
element?
*/

const max = 40

func part2() {
	data := buildData()
	// each pair produces a new letter to count
	allCounts := map[string][]map[rune]int{}
	for j := 0; j < len(data.curChain)-1; j++ {
		start := time.Now()
		key := data.curChain[j : j+2]
		inner(0, key, data.rules, allCounts)
		fmt.Println(j/2, key, time.Since(start))
	}
	// sum up all the counts for all the pairs in the top level
	counts := map[rune]int{}
	for j := 0; j < len(data.curChain)-1; j++ {
		key := data.curChain[j : j+2]
		for k2, v2 := range allCounts[key][0] {
			counts[k2] += v2
		}
	}
	// add in the counts for the initial string
	for _, v := range data.curChain {
		counts[v]++
	}
	minCount := math.MaxInt
	maxCount := 0
	for _, v := range counts {
		if v > maxCount {
			maxCount = v
		}
		if v < minCount {
			minCount = v
		}
	}
	fmt.Println(maxCount, minCount, maxCount-minCount)
}

func inner(depth int, pair string, rules map[string]rune, counts map[string][]map[rune]int) {
	// do we already know the answer for this pair at this depth?
	keyCounts, ok := counts[pair]
	if !ok {
		// no row for this pair yet -- make it!
		keyCounts = make([]map[rune]int, max)
		counts[pair] = keyCounts
	}
	// we have calculated this already
	if keyCounts[depth] != nil {
		return
	}
	// add on for my characters
	curMap := map[rune]int{}
	val := rules[pair]
	curMap[val]++
	if depth == max-1 {
		keyCounts[depth] = curMap
		return
	}
	// have we calculated the children?
	// if we haven't calculated them and sum them up and store them
	next1 := string([]byte{pair[0], byte(val)})
	nextCounts, ok := counts[next1]
	if len(nextCounts) == 0 || nextCounts[depth+1] == nil {
		inner(depth+1, next1, rules, counts)
		nextCounts = counts[next1] // reload
	}
	for k, v := range nextCounts[depth+1] {
		curMap[k] += v
	}
	next2 := string([]byte{byte(val), pair[1]})
	nextCounts2, ok := counts[next2]
	if len(nextCounts2) == 0 || nextCounts2[depth+1] == nil {
		inner(depth+1, next2, rules, counts)
		nextCounts2 = counts[next2] // reload
	}
	for k, v := range nextCounts2[depth+1] {
		curMap[k] += v
	}
	keyCounts[depth] = curMap
}

type Data struct {
	curChain string
	rules    map[string]rune
}

func buildData() Data {
	f, err := os.Open("./day14/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	//	scanner = bufio.NewScanner(strings.NewReader(`NNCB
	//
	//CH -> B
	//HH -> N
	//CB -> H
	//NH -> C
	//HB -> C
	//HC -> B
	//HN -> C
	//NN -> C
	//BH -> H
	//NC -> B
	//NB -> B
	//BN -> B
	//BB -> N
	//BC -> B
	//CC -> N
	//CN -> C`))
	scanner.Split(bufio.ScanLines)

	d := Data{
		rules: map[string]rune{},
	}
	scanner.Scan()
	d.curChain = scanner.Text()
	scanner.Scan() // blank line
	for scanner.Scan() {
		curRow := scanner.Text()
		parts := strings.Split(curRow, " -> ")
		d.rules[parts[0]] = rune(parts[1][0])
	}

	return d
}
