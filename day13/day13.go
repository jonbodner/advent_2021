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
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
The first section is a list of dots on the transparent paper. 0,0 represents the top-left coordinate. The first value, x, increases to the right. The second value, y, increases downward. So, the coordinate 3,0 is to the right of 0,0, and the coordinate 0,7 is below 0,0. The coordinates in this example form the following pattern, where # is a dot on the paper and . is an empty, unmarked position:

...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
Then, there is a list of fold instructions. Each instruction indicates a line on the transparent paper and wants you to fold the paper up (for horizontal y=... lines) or left (for vertical x=... lines). In this example, the first fold instruction is fold along y=7, which designates the line formed by all of the positions where y is 7 (marked here with -):

...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
-----------
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
Because this is a horizontal line, fold the bottom half up. Some of the dots might end up overlapping after the fold is complete, but dots will never appear exactly on a fold line. The result of doing this fold looks like this:

#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........
Now, only 17 dots are visible.

Notice, for example, the two dots in the bottom left corner before the transparent paper is folded; after the fold is complete, those dots appear in the top left corner (at 0,0 and 0,1). Because the paper is transparent, the dot just below them in the result (at 0,3) remains visible, as it can be seen through the transparent paper.

Also notice that some dots can end up overlapping; in this case, the dots merge together and become a single dot.

The second fold instruction is fold along x=5, which indicates this line:

#.##.|#..#.
#...#|.....
.....|#...#
#...#|.....
.#.#.|#.###
.....|.....
.....|.....
Because this is a vertical line, fold left:

#####
#...#
#...#
#...#
#####
.....
.....
The instructions made a square!

The transparent paper is pretty big, so for now, focus on just completing the first fold. After the first fold in the example above, 17 dots are visible - dots that end up overlapping after the fold is completed count as a single dot.

How many dots are visible after completing just the first fold instruction on your transparent paper?
*/
type Part1 struct {
	grid  [][]bool
	folds []Fold
}

type Fold struct {
	axis rune
	pos  int
}

func (p *Part1) Process(s string) {
	if len(strings.TrimSpace(s)) == 0 {
		return
	}
	if strings.HasPrefix(s, "fold along ") {
		info := strings.Split(s[11:], "=")
		pos, _ := strconv.Atoi(info[1])
		fold := Fold{
			axis: rune(info[0][0]),
			pos:  pos,
		}
		p.folds = append(p.folds, fold)
		return
	}
	xy := strings.Split(s, ",")
	y, _ := strconv.Atoi(xy[1])
	for i := len(p.grid); i <= y; i++ {
		p.grid = append(p.grid, []bool{})
	}
	x, _ := strconv.Atoi(xy[0])
	for i := len(p.grid[y]); i <= x; i++ {
		p.grid[y] = append(p.grid[y], false)
	}
	p.grid[y][x] = true
}

func (p *Part1) Result() int {
	firstFold := p.folds[0]
	switch firstFold.axis {
	case 'x':
		for i := 0; i < len(p.grid); i++ {
			for j := firstFold.pos + 1; j < len(p.grid[i]); j++ {
				if p.grid[i][j] {
					p.grid[i][firstFold.pos-(j-firstFold.pos)] = true
					p.grid[i][j] = false
				}
			}
		}
		return p.count(firstFold.pos, len(p.grid))
	case 'y':
	}
	return 0
}

func (p *Part1) count(maxX int, maxY int) int {
	total := 0
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX && j < len(p.grid[i]); j++ {
			if p.grid[i][j] {
				total++
			}
		}
	}
	return total
}

/*
Finish folding the transparent paper according to the instructions. The manual says the code is always eight capital letters.

What code do you use to activate the infrared thermal imaging camera system?
*/
type Part2 struct {
	grid  [][]bool
	folds []Fold
}

func (p *Part2) Process(s string) {
	if len(strings.TrimSpace(s)) == 0 {
		return
	}
	if strings.HasPrefix(s, "fold along ") {
		info := strings.Split(s[11:], "=")
		pos, _ := strconv.Atoi(info[1])
		fold := Fold{
			axis: rune(info[0][0]),
			pos:  pos,
		}
		p.folds = append(p.folds, fold)
		return
	}
	xy := strings.Split(s, ",")
	y, _ := strconv.Atoi(xy[1])
	for i := len(p.grid); i <= y; i++ {
		p.grid = append(p.grid, []bool{})
	}
	x, _ := strconv.Atoi(xy[0])
	for i := len(p.grid[y]); i <= x; i++ {
		p.grid[y] = append(p.grid[y], false)
	}
	p.grid[y][x] = true
}

func (p *Part2) Result() int {
	var lastX, lastY int
	for _, fold := range p.folds {
		switch fold.axis {
		case 'x':
			lastX = fold.pos
			for i := 0; i < len(p.grid); i++ {
				for j := fold.pos + 1; j < len(p.grid[i]); j++ {
					if p.grid[i][j] {
						p.grid[i][fold.pos-(j-fold.pos)] = true
						p.grid[i][j] = false
					}
				}
			}
		case 'y':
			lastY = fold.pos
			for i := fold.pos + 1; i < len(p.grid); i++ {
				for j := 0; j < len(p.grid[i]); j++ {
					if p.grid[i][j] {
						matchRow := fold.pos - (i - fold.pos)
						if len(p.grid[matchRow]) < j {
							for k := len(p.grid[matchRow]); k <= j; k++ {
								p.grid[matchRow] = append(p.grid[matchRow], false)
							}
						}
						p.grid[matchRow][j] = true
						p.grid[i][j] = false
					}
				}
			}
		}
	}
	p.printGrid(lastX, lastY)
	return 0
}

func (p *Part2) printGrid(maxX int, maxY int) {
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if p.grid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Processor interface {
	Process(s string)
	Result() int
}

func process(p Processor) {
	f, err := os.Open("./day13/input.txt")
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
