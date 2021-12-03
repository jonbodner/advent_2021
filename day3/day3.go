package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	process(&Part1{})
	process(&Part2{})
}

/*
You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the gamma rate and the epsilon rate). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:

00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate is 1.

The most common second bit of the numbers in the diagnostic report is 0, so the second bit of the gamma rate is 0.

The most common value of the third, fourth, and fifth bits are 1, 1, and 0, respectively, and so the final three bits of the gamma rate are 110.

So, the gamma rate is the binary number 10110, or 22 in decimal.

The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.
*/
type Part1 struct {
	total     int
	onesCount []int
}

func (p *Part1) Process(s string) {
	if p.onesCount == nil {
		p.onesCount = make([]int, len(s))
	}
	for i, b := range s {
		if b == '1' {
			p.onesCount[i]++
		}
	}
	p.total++
}

func (p *Part1) Result() int {
	var gamma int
	var epsilon int
	for _, v := range p.onesCount {
		gamma = gamma * 2
		epsilon = epsilon * 2
		if v > p.total/2 {
			gamma++
		} else {
			epsilon++
		}
	}
	return gamma * epsilon
}

/*
Both the oxygen generator rating and the CO2 scrubber rating are values that can be found in your diagnostic report - finding them is the tricky part. Both values are located using a similar process that involves filtering out values until only one remains. Before searching for either rating value, start with the full list of binary numbers from your diagnostic report and consider just the first bit of those numbers. Then:

Keep only numbers selected by the bit criteria for the type of rating value for which you are searching. Discard numbers which do not match the bit criteria.
If you only have one number left, stop; this is the rating value for which you are searching.
Otherwise, repeat the process, considering the next bit to the right.
The bit criteria depends on which type of rating value you want to find:

To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position,
and keep only numbers with that bit in that position.
If 0 and 1 are equally common, keep values with a 1 in the position being considered.

To find CO2 scrubber rating, determine the least common value (0 or 1) in the current bit position,
and keep only numbers with that bit in that position.
If 0 and 1 are equally common, keep values with a 0 in the position being considered.
*/
type Part2 struct {
	bits []string
}

func (p *Part2) Process(s string) {
	p.bits = append(p.bits, s)
}

func (p *Part2) Result() int {
	o2 := p.Find('1', '0')
	co2 := p.Find('0', '1')
	o2Level, _ := strconv.ParseInt(o2, 2, 64)
	co2Level, _ := strconv.ParseInt(co2, 2, 64)
	return int(o2Level * co2Level)
}

func (p *Part2) Find(gt byte, lt byte) string {
	o2s := p.bits
	pos := 0
	for len(o2s) > 1 && pos < len(o2s[0]) {
		onesCount := buildOnesCount(o2s, pos)
		var o2Check = gt
		if onesCount*2 < len(o2s) {
			o2Check = lt
		}
		// keep entries from o2s with a o2Check in the ith position
		o2s = filter(pos, o2Check, o2s)
		pos++
	}
	return o2s[0]
}

func buildOnesCount(s []string, pos int) int {
	var out int
	for _, v := range s {
		if v[pos] == '1' {
			out++
		}
	}
	return out
}

func filter(pos int, ch byte, s []string) []string {
	var newO2s []string
	for _, v := range s {
		if v[pos] == ch {
			newO2s = append(newO2s, v)
		}
	}
	return newO2s
}

type Processor interface {
	Process(s string)
	Result() int
}

func process(p Processor) {
	f, err := os.Open("./day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		p.Process(scanner.Text())
	}
	fmt.Println(p.Result())
}
