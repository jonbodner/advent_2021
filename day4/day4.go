package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

/*
Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?
*/

type board [][]string

func (b board) score(bs boardstate, lastNum int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !bs[i][j] {
				n, _ := strconv.Atoi(b[i][j])
				sum += n
			}
		}
	}
	return sum * lastNum
}

func (b board) contains(num string) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] == num {
				return i, j
			}
		}
	}
	return -1, -1
}

type boardstate [5][5]bool

func (bs boardstate) won() bool {
	// check rows
	for i := 0; i < 5; i++ {
		won := true
		for j := 0; j < 5; j++ {
			if !bs[i][j] {
				won = false
				break
			}
		}
		if won {
			fmt.Println("row", i)
			return true
		}
	}

	// check cols
	for j := 0; j < 5; j++ {
		won := true
		for i := 0; i < 5; i++ {
			if !bs[i][j] {
				won = false
				break
			}
		}
		if won {
			fmt.Println("col", j)
			return true
		}
	}
	return false
}

func part1() {
	numbers, boards := getData()
	//now track values in each board, see if it wins
	boardstates := make([]boardstate, len(boards))
	for _, v := range numbers {
		fmt.Println(v)
		for p, b := range boards {
			i, j := b.contains(v)
			if i != -1 {
				boardstates[p][i][j] = true
				if boardstates[p].won() {
					fmt.Println("winner!", b)
					lastNum, _ := strconv.Atoi(v)
					fmt.Println(b.score(boardstates[p], lastNum))
					return
				}
			}
		}
	}
}

func part2() {
	numbers, boards := getData()
	//now track values in each board, see if it wins
	boardstates := make([]boardstate, len(boards))
	didWin := make([]bool, len(boards))
	for _, v := range numbers {
		fmt.Println(v)
		for p, b := range boards {
			if didWin[p] {
				continue
			}
			i, j := b.contains(v)
			if i != -1 {
				boardstates[p][i][j] = true
				if boardstates[p].won() {
					fmt.Println("winner!", b)
					lastNum, _ := strconv.Atoi(v)
					fmt.Println(b.score(boardstates[p], lastNum))
					didWin[p] = true
				}
			}
		}
	}
}

func getData() ([]string, []board) {
	f, err := os.Open("./day4/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	//read calls
	scanner.Scan()
	numberLine := scanner.Text()
	numbers := strings.Split(numberLine, ",")
	var boards []board
	//read boards
	for scanner.Scan() {
		//skip blank line
		var curBoard board
		//lines 1- 5
		for i := 0; i < 5; i++ {
			scanner.Scan()
			vals := strings.Fields(scanner.Text())
			curBoard = append(curBoard, vals)
		}
		boards = append(boards, curBoard)
	}
	f.Close()
	return numbers, boards
}
