package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	process(&Part1{})
	process(&Part2{})
}

/*
Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where x1,y1 are the coordinates of one end the line segment and x2,y2 are the coordinates of the other end. These line segments include the points at both ends. In other words:

An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.

So, the horizontal and vertical lines from the above list would produce the following diagram:

.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9. Each position is shown as the number of lines which cover that point or . if no line covers that point. The top-left pair of 1s, for example, comes from 2,2 -> 2,1; the very bottom row is formed by the overlapping lines 0,9 -> 5,9 and 0,9 -> 2,9.

To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.

Consider only horizontal and vertical lines. At how many points do at least two lines overlap?
*/
type Part1 struct {
	board [][]int
}

func (p *Part1) Process(s string) {
	startParts, endParts := parseStartEnd(s)
	// fill in y
	if startParts[0] == endParts[0] {
		x, _ := strconv.Atoi(startParts[0])
		startY, _ := strconv.Atoi(startParts[1])
		endY, _ := strconv.Atoi(endParts[1])
		//swap if order wrong
		if startY > endY {
			temp := startY
			startY = endY
			endY = temp
		}
		// grow if needed
		for i := len(p.board); i <= endY; i++ {
			p.board = append(p.board, []int{})
		}
		for i := startY; i <= endY; i++ {
			//grow if needed
			rowLen := len(p.board[i])
			for j := rowLen; j <= x; j++ {
				p.board[i] = append(p.board[i], 0)
			}
			p.board[i][x]++
		}
	} else if startParts[1] == endParts[1] {
		y, _ := strconv.Atoi(startParts[1])
		startX, _ := strconv.Atoi(startParts[0])
		endX, _ := strconv.Atoi(endParts[0])
		//swap if order wrong
		if startX > endX {
			temp := startX
			startX = endX
			endX = temp
		}
		// grow if needed
		for i := len(p.board); i <= y; i++ {
			p.board = append(p.board, []int{})
		}
		//grow if needed
		rowLen := len(p.board[y])
		for j := rowLen; j <= endX; j++ {
			p.board[y] = append(p.board[y], 0)
		}
		for i := startX; i <= endX; i++ {
			p.board[y][i]++
		}
	} else {
		fmt.Println("skip, diagonal: ", s)
	}
}

func parseStartEnd(s string) ([]string, []string) {
	// 781,721 -> 781,611
	parts := strings.Split(s, "->")
	startParts := strings.Split(parts[0], ",")
	endParts := strings.Split(parts[1], ",")
	startParts[0] = strings.TrimSpace(startParts[0])
	startParts[1] = strings.TrimSpace(startParts[1])
	endParts[0] = strings.TrimSpace(endParts[0])
	endParts[1] = strings.TrimSpace(endParts[1])
	return startParts, endParts
}

func (p *Part1) Result() int {
	count := 0
	for _, v := range p.board {
		for _, w := range v {
			if w > 1 {
				count++
			}
		}
	}
	return count
}

/*
Unfortunately, considering only horizontal and vertical lines doesn't give you the full picture; you need to also consider diagonal lines.

Because of the limits of the hydrothermal vent mapping system, the lines in your list will only ever be horizontal, vertical, or a diagonal line at exactly 45 degrees. In other words:

An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.
Considering all lines from the above example would now produce the following diagram:

1.1....11.
.111...2..
..2.1.111.
...1.2.2..
.112313211
...1.2....
..1...1...
.1.....1..
1.......1.
222111....
You still need to determine the number of points where at least two lines overlap. In the above example, this is still anywhere in the diagram with a 2 or larger - now a total of 12 points.
*/
type Part2 struct {
	board [][]int
}

func (p *Part2) Process(s string) {
	startParts, endParts := parseStartEnd(s)
	// fill in y
	if startParts[0] == endParts[0] {
		x, _ := strconv.Atoi(startParts[0])
		startY, _ := strconv.Atoi(startParts[1])
		endY, _ := strconv.Atoi(endParts[1])
		//swap if order wrong
		if startY > endY {
			temp := startY
			startY = endY
			endY = temp
		}
		// grow if needed
		for i := len(p.board); i <= endY; i++ {
			p.board = append(p.board, []int{})
		}
		for i := startY; i <= endY; i++ {
			//grow if needed
			rowLen := len(p.board[i])
			for j := rowLen; j <= x; j++ {
				p.board[i] = append(p.board[i], 0)
			}
			p.board[i][x]++
		}
	} else if startParts[1] == endParts[1] {
		y, _ := strconv.Atoi(startParts[1])
		startX, _ := strconv.Atoi(startParts[0])
		endX, _ := strconv.Atoi(endParts[0])
		//swap if order wrong
		if startX > endX {
			temp := startX
			startX = endX
			endX = temp
		}
		// grow if needed
		for i := len(p.board); i <= y; i++ {
			p.board = append(p.board, []int{})
		}
		//grow if needed
		rowLen := len(p.board[y])
		for j := rowLen; j <= endX; j++ {
			p.board[y] = append(p.board[y], 0)
		}
		for i := startX; i <= endX; i++ {
			p.board[y][i]++
		}
	} else {
		startY, _ := strconv.Atoi(startParts[1])
		endY, _ := strconv.Atoi(endParts[1])
		startX, _ := strconv.Atoi(startParts[0])
		endX, _ := strconv.Atoi(endParts[0])
		if startY > endY {
			// swap so we are always working down
			tempX := startX
			tempY := startY
			startX = endX
			startY = endY
			endX = tempX
			endY = tempY
		}
		// make sure we have enough rows
		curLen := len(p.board)
		for i := curLen; i <= endY; i++ {
			p.board = append(p.board, []int{})
		}
		incX := 1
		if startX > endX {
			incX = -1
		}
		curX := startX
		for y := startY; y <= endY; y++ {
			curRowLen := len(p.board[y])
			for x := curRowLen; x <= curX; x++ {
				p.board[y] = append(p.board[y], 0)
			}
			p.board[y][curX]++
			curX += incX
		}
	}
}

func (p *Part2) Result() int {
	count := 0
	for _, v := range p.board {
		for _, w := range v {
			if w > 1 {
				count++
			}
		}
	}
	return count
}

type Processor interface {
	Process(s string)
	Result() int
}

func process(p Processor) {
	f, err := os.Open("./day5/input.txt")
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
