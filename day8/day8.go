package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	process(&Part1{})
	process(&Part2{})
}

/*
Each digit of a seven-segment display is rendered by turning on or off any of seven segments named a through g:

  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
So, to render a 1, only segments c and f would be turned on; the rest would be off. To render a 7, only segments a, c, and f would be turned on.

The problem is that the signals which control the segments have been mixed up on each display. The submarine is still
trying to display numbers by producing output on signal wires a through g, but those wires are connected to segments
randomly. Worse, the wire/segment connections are mixed up separately for each four-digit display! (All of the digits
within a display use the same connections, though.)

So, you might know that only signal wires b and g are turned on, but that doesn't mean segments b and g are turned on:
the only digit that uses two segments is 1, so it must mean segments c and f are meant to be on. With just that
information, you still can't tell which wire (b/g) goes to which segment (c/f). For that, you'll need to collect
more information.

For each display, you watch the changing signals for a while, make a note of all ten unique signal patterns you see,
and then write down a single four digit output value (your puzzle input). Using the signal patterns, you should be
able to work out which pattern corresponds to which digit.

For example, here is what you might see in a single entry in your notes:

acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf
(The entry is wrapped here to two lines so it fits; in your notes, it will all be on a single line.)

Each entry consists of ten unique signal patterns, a | delimiter, and finally the four digit output value. Within an
entry, the same wire/segment connections are used (but you don't know what the connections actually are). The unique
signal patterns correspond to the ten different ways the submarine tries to render a digit using the current wire/segment
connections. Because 7 is the only digit that uses three segments, dab in the above example means that to render a 7,
signal lines d, a, and b are on. Because 4 is the only digit that uses four segments, eafb means that to render a 4,
signal lines e, a, f, and b are on.

Using this information, you should be able to work out which combination of signal wires corresponds to each of the ten
digits. Then, you can decode the four digit output value. Unfortunately, in the above example, all of the digits in the
output value (cdfeb fcadb cdfeb cdbaf) use five segments and are more difficult to deduce.

For now, focus on the easy digits. Consider this larger example:

be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce

Because the digits 1, 4, 7, and 8 each use a unique number of segments, you should be able to tell which combinations of signals correspond to those digits.
Counting only digits in the output values (the part after | on each line), in the above example, there are 26 instances of digits that use a unique number
of segments (highlighted above).

In the output values, how many times do digits 1, 4, 7, or 8 appear?
*/
type Part1 struct {
	counter int
}

func (p *Part1) Process(s string) {
	// throw away everything before the |
	parts := strings.Split(s, "|")
	parts2 := strings.Fields(parts[1])
	for _, v := range parts2 {
		switch len(v) {
		case 2, 4, 3, 7:
			p.counter++

		}
	}
}

func (p *Part1) Result() int {
	return p.counter
}

/*
After some careful analysis, the mapping between signal wires and segments only make sense in the following configuration:

 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc
So, the unique signal patterns would correspond to the following digits:

acedgfb: 8
cdfbe: 5
gcdfa: 2
fbcad: 3
dab: 7
cefabd: 9
cdfgeb: 6
eafb: 4
cagedb: 0
ab: 1
Then, the four digits of the output value can be decoded:

cdfeb: 5
fcadb: 3
cdfeb: 5
cdbaf: 3
Therefore, the output value for this entry is 5353.

Following this same process for each entry in the second, larger example above, the output value of each entry can be determined:

fdgacbe cefdb cefbgd gcbe: 8394
fcgedb cgb dgebacf gc: 9781
cg cg fdcagb cbg: 1197
efabcd cedba gadfec cb: 9361
gecf egdcabf bgf bfgea: 4873
gebdcfa ecba ca fadegcb: 8418
cefg dcbef fcge gbcadfe: 4548
ed bcgafe cdgba cbgef: 1625
gbdfcae bgc cg cgb: 8717
fgae cfgab fg bagce: 4315
Adding all of the output values in this larger example produces 61229.

For each entry, determine all of the wire/segment connections and decode the four-digit output values. What do you get if you add up all of the output values?
*/
type Part2 struct {
	total int
}

func (p *Part2) Process(s string) {
	possible := map[int][]string{}
	parts := strings.Split(s, "|")
	nums := strings.Fields(parts[0])
	for _, v := range nums {
		possible[len(v)] = append(possible[len(v)], v)
	}
	fmt.Println(possible)
	association := map[string]string{}
	// item that's in 7, but not in 1 is a
	// 1 is len 2, 7 is len 3
	association["a"] = returnDiff(possible[2][0], possible[3][0])
	fmt.Println("a is", association["a"])
	// other fields in 4 are b and d
	// 1 is len 2, 4 is len 4
	bd := returnDiff(possible[2][0], possible[4][0])
	fmt.Println(bd)
	// if it's in all of length 6, it's b
	// if it's in 2 of length 6, it's d
	for _, v := range possible[6] {
		maybeD := returnDiff(v, bd)
		if maybeD != "" {
			association["d"] = maybeD
			if maybeD == bd[:1] {
				association["b"] = bd[1:]
			} else {
				association["b"] = bd[:1]
			}
			break
		}
	}
	fmt.Println("b is", association["b"])
	fmt.Println("d is", association["d"])

	// 5 is the one of length 5 with 3 known
	known := association["a"] + association["b"] + association["d"]
	for _, v := range possible[5] {
		remaining := returnDiff(known, v)
		if len(remaining) == 2 {
			fmt.Println(remaining)
			// the one in 5 but not in 1 is g
			association["g"] = returnDiff(possible[2][0], remaining)
			fmt.Println("g is", association["g"])
			// the one in 1 but not in 5 is c
			association["c"] = returnDiff(remaining, possible[2][0])
			fmt.Println("c is", association["c"])
			// the unknown one in 5 and 1 is f
			known = known + association["g"] + association["c"]
			association["f"] = returnDiff(known, remaining)
			fmt.Println("f is", association["f"])
			known = known + association["f"]
			break
		}
	}
	// e is whatever letter isn't mapped yet
	association["e"] = returnDiff(known, "abcdefg")
	fmt.Println("e is", association["e"])
	// invert the map
	invert := map[string]string{}
	for k, v := range association {
		invert[v] = k
	}

	/*
		1 is fg
		7 is dfg
		d == a
		f is c or f
		g is c or f
		4 is fcgb
		c is b or d
		b is b or d
		c is all of length 6 so b
		b is in 2 of length 6 so d
		c == b
		b == d
		5 is the one with 3 known == bdcfa (dabxx)
		look at 5, know a, b, d, don't know f or g
		the one in 1 is f
		f == f
		the one not in 1 is g
		a == g
		the one in 1 not in 5 is c
		g == c
		the one not figured out is e
		e == e
	*/

	parts2 := strings.Fields(parts[1])
	fmt.Println(parts2)
	fmt.Println(invert)
	number := 0
	for _, v := range parts2 {
		actual := convert(v, invert)
		number = number*10 + translate(actual)
	}
	fmt.Println(number)
	p.total += number
}

func convert(in string, association map[string]string) string {
	var out string
	for _, v := range in {
		out = out + association[string(v)]
	}
	return out
}

var numberParts = []string{
	"abcefg",  // 0
	"cf",      // 1
	"acdeg",   // 2
	"acdfg",   // 3
	"bcdf",    // 4
	"abdfg",   // 5
	"abdefg",  // 6
	"acf",     // 7
	"abcdefg", // 8
	"abcdfg",  // 9
}

func translate(actual string) int {
	b := []byte(actual)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	s := string(b)
	for i, v := range numberParts {
		if s == v {
			return i
		}
	}
	panic("should never happen: " + actual + " " + s)
}

func returnDiff(a, b string) string {
	am := map[rune]bool{}
	for _, v := range a {
		am[v] = true
	}
	var out string
	for _, v := range b {
		if !am[v] {
			out = out + string(v)
		}
	}
	return out
}

func (p *Part2) Result() int {
	return p.total
}

type Processor interface {
	Process(s string)
	Result() int
}

func process(p Processor) {
	f, err := os.Open("./day8/input.txt")
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
