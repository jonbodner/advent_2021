package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	process(&Part1{
		nodes: map[string]*Node{},
	})
	process(&Part2{
		nodes: map[string]*Node{},
	})
}

/*
Fortunately, the sensors are still mostly working, and so you build a rough map of the remaining caves (your puzzle input). For example:

start-A
start-b
A-c
A-b
b-d
A-end
b-end
This is a list of how all of the caves are connected. You start in the cave named start, and your destination is the cave named end. An entry like b-d means that cave b is connected to cave d - that is, you can move between them.

So, the above cave system looks roughly like this:

    start
    /   \
c--A-----b--d
    \   /
     end
Your goal is to find the number of distinct paths that start at start, end at end, and don't visit small caves more than once. There are two types of caves: big caves (written in uppercase, like A) and small caves (written in lowercase, like b). It would be a waste of time to visit any small cave more than once, but big caves are large enough that it might be worth visiting them multiple times. So, all paths you find should visit small caves at most once, and can visit big caves any number of times.

Given these rules, there are 10 paths through this example cave system:

start,A,b,A,c,A,end
start,A,b,A,end
start,A,b,end
start,A,c,A,b,A,end
start,A,c,A,b,end
start,A,c,A,end
start,A,end
start,b,A,c,A,end
start,b,A,end
start,b,end
(Each line in the above list corresponds to a single path; the caves visited by that path are listed in the order they are visited and separated by commas.)

Note that in this cave system, cave d is never visited by any path: to do so, cave b would need to be visited twice (once on the way to cave d and a second time when returning from cave d), and since cave b is small, this is not allowed.

How many paths through this cave system are there that visit small caves at most once?
*/
type Part1 struct {
	nodes     map[string]*Node
	startNode *Node
}

func (p *Part1) Process(s string) {
	if len(strings.TrimSpace(s)) == 0 {
		return
	}
	// bm-XY
	nodes := strings.Split(s, "-")
	node1Name := nodes[0]
	node2Name := nodes[1]
	node1 := p.buildFindNode(node1Name)
	node2 := p.buildFindNode(node2Name)
	node1.connections = append(node1.connections, node2)
	node2.connections = append(node2.connections, node1)
}

func (p *Part1) buildFindNode(nodeName string) *Node {
	if node, ok := p.nodes[nodeName]; !ok {
		node = &Node{
			name: nodeName,
			big:  unicode.IsUpper(rune(nodeName[0])),
		}
		if p.nodes == nil {
			p.nodes = map[string]*Node{}
		}
		p.nodes[nodeName] = node
		if nodeName == "start" {
			p.startNode = node
		}
		return node
	} else {
		return node
	}
}

func (p *Part1) Result() int {
	start := time.Now()
	total := findPaths(p.startNode, []*Node{p.startNode})
	fmt.Println(time.Since(start))
	return total
}

/*
After reviewing the available paths, you realize you might have time to visit a single small cave twice.
Specifically, big caves can be visited any number of times, a single small cave can be visited at most twice,
and the remaining small caves can be visited at most once. However, the caves named start and end can only be
visited exactly once each: once you leave the start cave, you may not return to it, and once you reach the end
cave, the path must end immediately.

Given these new rules, how many paths through this cave system are there?
*/
type Part2 struct {
	nodes     map[string]*Node
	startNode *Node
}

func (p *Part2) Process(s string) {
	if len(strings.TrimSpace(s)) == 0 {
		return
	}
	// bm-XY
	nodes := strings.Split(s, "-")
	node1Name := nodes[0]
	node2Name := nodes[1]
	node1 := p.buildFindNode(node1Name)
	node2 := p.buildFindNode(node2Name)
	node1.connections = append(node1.connections, node2)
	node2.connections = append(node2.connections, node1)
}

func (p *Part2) buildFindNode(nodeName string) *Node {
	if node, ok := p.nodes[nodeName]; ok {
		return node
	}
	node := &Node{
		name: nodeName,
		big:  unicode.IsUpper(rune(nodeName[0])),
	}
	p.nodes[nodeName] = node
	if nodeName == "start" {
		p.startNode = node
	}
	return node
}

func (p *Part2) Result() int {
	start := time.Now()
	total := findPaths2(p.startNode, []*Node{p.startNode}, false)
	fmt.Println(time.Since(start))
	return total
}

func findPaths(node *Node, path []*Node) int {
	total := 0
outer:
	for _, v := range node.connections {
		if v.name == "end" {
			//printPath(path)
			total++
			continue
		}
		if !v.big {
			for _, p := range path {
				if p.name == v.name {
					continue outer
				}
			}
		}
		newPath := make([]*Node, len(path), len(path)+10)
		copy(newPath, path)
		newPath = append(newPath, v)
		total += findPaths(v, newPath)
	}
	return total
}

func findPaths2(node *Node, path []*Node, doubleSmall bool) int {
	total := 0
outer:
	for _, v := range node.connections {
		doubleSmall := doubleSmall
		if v.name == "end" {
			//printPath(path)
			total++
			continue
		}
		if v.name == "start" {
			continue outer
		}
		if !v.big {
			for _, p := range path {
				if p.name == v.name {
					if doubleSmall {
						continue outer
					}
					doubleSmall = true
				}
			}
		}
		newPath := make([]*Node, len(path), len(path)+1)
		copy(newPath, path)
		newPath = append(newPath, v)
		total += findPaths2(v, newPath, doubleSmall)
	}
	return total
}

func printPath(path []*Node) {
	for _, v := range path {
		fmt.Print(v.name, ",")
	}
	fmt.Println("end")
}

type Node struct {
	name        string
	big         bool
	connections []*Node
}

type Processor interface {
	Process(s string)
	Result() int
}

func process(p Processor) {
	f, err := os.Open("./day12/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	//	scanner = bufio.NewScanner(strings.NewReader(`
	//start-A
	//start-b
	//A-c
	//A-b
	//b-d
	//A-end
	//b-end`))

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		p.Process(scanner.Text())
	}
	fmt.Println(p.Result())
}
