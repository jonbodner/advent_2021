package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	part1()
	part2()
}

/*
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
You start in the top left position, your destination is the bottom right position, and you cannot move diagonally.
The number at each position is its risk level; to determine the total risk of an entire path, add up the risk levels
of each position you enter (that is, don't count the risk level of your starting position unless you enter it; leaving
it adds no risk to your total).

The total risk of this path is 40 (the starting position is never entered, so its risk is not counted).

What is the lowest total risk of any path from the top left to the bottom right?
*/
func part1() {
	g := loadData()
	start := point{0, 0}
	st := time.Now()
	dist, _ := dijkstra(g, start)
	fmt.Println(time.Since(st))
	fmt.Println(dist[point{len(g) - 1, len(g) - 1}])
}

/*
The entire cave is actually five times larger in both dimensions than you thought; the area you originally scanned
is just one tile in a 5x5 tile area that forms the full map. Your original map tile repeats to the right and downward;
each time the tile repeats to the right or downward, all of its risk levels are 1 higher than the tile immediately
up or left of it. However, risk levels above 9 wrap back around to 1. So, if your original map had some position with
a risk level of 8, then that same position on each of the 25 total tiles would be as follows:

8 9 1 2 3
9 1 2 3 4
1 2 3 4 5
2 3 4 5 6
3 4 5 6 7
Each single digit above corresponds to the example position with a value of 8 on the top-left tile. Because the full
map is actually five times larger in both dimensions, that position appears a total of 25 times, once in each duplicated
tile, with the values shown above.

The total risk of this path is 315 (the starting position is still never entered, so its risk is not counted).

Using the full map, what is the lowest total risk of any path from the top left to the bottom right?
*/
func part2() {
	g := loadData()
	gg := growData(g)
	//printGrid(gg)
	start := point{0, 0}
	st := time.Now()
	dist, _ := dijkstra(gg, start)
	fmt.Println(time.Since(st))
	fmt.Println(dist[point{len(gg) - 1, len(gg) - 1}])
}

func growData(g [][]byte) [][]byte {
	out := make([][]byte, len(g)*5)
	for i := 0; i < len(g)*5; i++ {
		out[i] = make([]byte, len(g)*5)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < len(g); k++ {
				for m := 0; m < len(g); m++ {
					newVal := g[k][m] + byte(i+j)
					if newVal > 9 {
						newVal = newVal - 9
					}
					//fmt.Println(k+i*len(g), m+j*len(g), newVal)
					out[k+i*len(g)][m+j*len(g)] = newVal
				}
			}
		}
	}
	return out
}

type point struct {
	x, y int
}

func printGrid(g [][]byte) {
	for _, v := range g {
		for _, c := range v {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

/*
In the following pseudocode algorithm, dist is an array that contains the current distances from the source to
other vertices, i.e. dist[u] is the current distance from the source to the vertex u. The prev array contains pointers
to previous-hop nodes on the shortest path from source to the given vertex (equivalently, it is the next-hop on the
path from the given vertex to the source). The code u ← vertex in Q with min dist[u], searches for the vertex u in the
vertex set Q that has the least dist[u] value. length(u, v) returns the length of the edge joining (i.e. the distance
between) the two neighbor-nodes u and v. The variable alt on line 18 is the length of the path from the root node to
the neighbor node v if it were to go through u. If this path is shorter than the current shortest path recorded for v,
that current path is replaced with this alt path.

 1  function Dijkstra(Graph, source):
 2
 3      create vertex set Q
 4
 5      for each vertex v in Graph:
 6          dist[v] ← INFINITY
 7          prev[v] ← UNDEFINED
 8          add v to Q
 9      dist[source] ← 0
10
11      while Q is not empty:
12          u ← vertex in Q with min dist[u]
13
14          remove u from Q
15
16          for each neighbor v of u still in Q:
17              alt ← dist[u] + length(u, v)
18              if alt < dist[v]:
19                  dist[v] ← alt
20                  prev[v] ← u
21
22      return dist[], prev[]
*/
func dijkstra(graph [][]byte, source point) (map[point]int, map[point]point) {
	q := map[point]bool{}
	dist := map[point]int{}
	prev := map[point]point{}
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			q[point{j, i}] = true
		}
	}
	dist[source] = 0
	outDist := map[point]int{}
	for len(q) > 0 {
		u, distU := minDistance(dist)
		delete(q, u)
		delete(dist, u)
		outDist[u] = distU
		n := neighbors(u, q)
		for _, v := range n {
			alt := distU + int(graph[v.y][v.x])
			curDist, ok := dist[v]
			if !ok {
				curDist = math.MaxInt
			}
			if alt < curDist {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	return outDist, prev
}

func minDistance(points map[point]int) (point, int) {
	var lowest point
	lowestScore := math.MaxInt
	for p, score := range points {
		if score < lowestScore {
			lowestScore = score
			lowest = p
		}
	}
	return lowest, lowestScore
}

func neighbors(u point, q map[point]bool) []point {
	var out []point
	if q[point{u.x, u.y - 1}] {
		out = append(out, point{u.x, u.y - 1})
	}
	if q[point{u.x, u.y + 1}] {
		out = append(out, point{u.x, u.y + 1})
	}
	if q[point{u.x - 1, u.y}] {
		out = append(out, point{u.x - 1, u.y})
	}
	if q[point{u.x + 1, u.y}] {
		out = append(out, point{u.x + 1, u.y})
	}
	return out
}

func loadData() [][]byte {
	contents, _ := os.ReadFile("./day15/input.txt")
	//	contents = []byte(`1163751742
	//1381373672
	//2136511328
	//3694931569
	//7463417111
	//1319128137
	//1359912421
	//3125421639
	//1293138521
	//2311944581`)
	grid := bytes.Split(contents, []byte{'\n'})
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = grid[i][j] - '0'
		}
	}
	if len(grid[len(grid)-1]) == 0 {
		grid = grid[:len(grid)-1]
	}
	// last line is blank
	return grid
}
