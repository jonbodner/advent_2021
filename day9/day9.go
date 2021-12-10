package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	part1()
	part2()
}

/*
Smoke flows to the lowest point of the area it's in. For example, consider the following heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678
Each number corresponds to the height of a particular location, where 9 is the highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent locations. Most locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1 and a 0), one is in the third row (a 5), and one is in the bottom row (also a 5). All other locations on the heightmap have some lower adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low points on your heightmap?
*/
func part1() {
	grid := getInitial()
	total := 0
	for y, row := range grid {
		for x, cell := range row {
			var vals []byte
			if y > 0 {
				vals = append(vals, grid[y-1][x])
			}
			if x > 0 {
				vals = append(vals, grid[y][x-1])
			}
			if y < len(grid)-1 {
				vals = append(vals, grid[y+1][x])
			}
			if x < len(row)-1 {
				vals = append(vals, grid[y][x+1])
			}
			low := true
			for _, v := range vals {
				if cell >= v {
					low = false
					break
				}
			}
			if low {
				//fmt.Println("it's the lowest!")
				//fmt.Printf("at (%d,%d), comparing %d to %v\n", x, y, cell, vals)
				// the vals here are ASCII characters, not numbers, so subtract '0'
				total += int(cell-'0') + 1
			}
		}
	}
	fmt.Println(total)
}

/*
A basin is all locations that eventually flow downward to a single low point. Therefore, every low point has a basin, although some basins are very small. Locations of height 9 do not count as being in any basin, and all other locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the low point. The example above has four basins.

The top-left basin, size 3:

2199943210
3987894921
9856789892
8767896789
9899965678
The top-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
The middle basin, size 14:

2199943210
3987894921
9856789892
8767896789
9899965678
The bottom-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
Find the three largest basins and multiply their sizes together. In the above example, this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins?
*/
func part2() {
	/*
		divide up the space by the 9s. A basin is an area surrounded by the edge and by 9s.
		surrounded only means UDLR.
		map coloring!
		scan from top to bottom, left to right
		start with 1, check up (if there's an up), if it's already colored, then set to same color
		increment color when hit a 9
		if using a color and find there's a contiguous color, go back and recolor everything of that color
		count the number of numbers (colors), take top 3, multiply
	*/
	grid := getInitial()
	var colors [][]int
	colorList := map[int]int{}
	curColor := 1
	in9 := false
	for y, row := range grid {
		colors = append(colors, []int{})
		for _, cell := range row {
			if cell == '9' {
				colors[y] = append(colors[y], -1)
				if !in9 {
					curColor++
				}
				in9 = true
				continue
			}
			in9 = false
			colors[y] = append(colors[y], curColor)
			colorList[curColor]++
		}
		// change color at end of row
		curColor++
	}

	printGrid(colors)
	fmt.Println(colorList)
	// unification pass, check the cell above, see if it's a different color,
	// if so change all of the same color to the color of the cell above
	for y := 1; y < len(colors); y++ {
		for x := 0; x < len(colors[y]); x++ {
			curColor := colors[y][x]
			// if we've already merged this color, skip
			if _, ok := colorList[curColor]; !ok {
				continue
			}
			upColor := colors[y-1][x]
			if curColor != -1 && upColor != -1 && curColor != upColor {
				colorList[upColor] += colorList[curColor]
				delete(colorList, curColor)
				for x1 := 0; x1 < len(colors[y]); x1++ {
					if colors[y][x1] == curColor {
						colors[y][x1] = upColor
					}
				}
			}
		}
	}
	printGrid(colors)
	fmt.Println(colorList)
	// now count and sort and multiply
	totals := make([]int, 0, len(colorList))
	for _, v := range colorList {
		totals = append(totals, v)
	}
	sort.Ints(totals)
	fmt.Println(totals)
	fmt.Println(totals[len(totals)-1] * totals[len(totals)-2] * totals[len(totals)-3])
}

func printGrid(colors [][]int) {
	curSymbol := '¡'
	for i := 0; i < len(colors); i++ {
		for j := 0; j < len(colors[i]); j++ {
			fmt.Print(string(curSymbol + rune(colors[i][j])))
		}
		fmt.Println()
	}
}

func getInitial() [][]byte {
	contents, _ := os.ReadFile("./day9/input.txt")
	grid := bytes.Split(contents, []byte{'\n'})
	if len(grid[len(grid)-1]) == 0 {
		grid = grid[:len(grid)-1]
	}
	// last line is blank
	return grid
}
