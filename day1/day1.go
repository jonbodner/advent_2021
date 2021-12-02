package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
The first order of business is to figure out how quickly the depth increases, just so you know what you're dealing with - you never know if the keys will get carried into deeper water by an ocean current or a fish or something.

To do this, count the number of times a depth measurement increases from the previous measurement. (There is no measurement before the first measurement.) In the example above, the changes are as follows:

199 (N/A - no previous measurement)
200 (increased)
208 (increased)
210 (increased)
200 (decreased)
207 (increased)
240 (increased)
269 (increased)
260 (decreased)
263 (increased)
In this example, there are 7 measurements that are larger than the previous measurement.

*/
func main() {
	process(&Part1{})
	process(&Part2{})
}

type Processor interface {
	Process(s string)
	Result() int
}

type Part1 struct {
	count int
	last  int
}

func (p *Part1) Process(s string) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if p.last > 0 && i > p.last {
		p.count++
	}
	p.last = i
}

func (p *Part1) Result() int {
	return p.count
}

func process(p Processor) {
	f, err := os.Open("./day1/input.txt")
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

type Part2 struct {
	count int
	last  [3]int
}

func (p *Part2) Process(s string) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if p.last[0] == 0 {
		p.last[0] = i
		return
	}
	if p.last[1] == 0 {
		p.last[1] = i
		return
	}
	if p.last[2] == 0 {
		p.last[2] = i
		return
	}
	curSum := p.last[0] + p.last[1] + p.last[2]
	if curSum < p.last[1]+p.last[2]+i {
		p.count++
	}
	p.last[0] = p.last[1]
	p.last[1] = p.last[2]
	p.last[2] = i
}

func (p *Part2) Result() int {
	return p.count
}
