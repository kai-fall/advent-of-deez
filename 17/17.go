package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	a    int
	prog []int
)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func mаp[A any, B any](s []A, f func(A) B) []B {
	out := make([]B, len(s))
	for i, v := range s {
		out[i] = f(v)
	}
	return out
}

func init() {
	content, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	a = atoi(strings.Split(lines[0], " ")[2])
	prog = mаp(strings.Split(strings.Split(lines[4], " ")[1], ","), atoi)
}

func run(a int) []int {
	b, c := 0, 0
	out := []int{}
	for p := 0; p < len(prog); p += 2 {
		opcode := prog[p]
		operand := prog[p+1]
		value := operand
		switch operand {
		case 4:
			value = a
		case 5:
			value = b
		case 6:
			value = c
		}
		switch opcode {
		case 0:
			a >>= value
		case 1:
			b ^= operand
		case 2:
			b = value % 8
		case 3:
			if a != 0 {
				p = operand - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, value%8)
		case 6:
			b = a >> value
		case 7:
			c = a >> value
		}
	}
	return out
}

func part1() {
	output := run(a)
	fmt.Println(strings.Join(mаp(output, strconv.Itoa), ","))
}

func part2() {
	a := 0
	for i := len(prog) - 1; i >= 0; i-- {
		a <<= 3
		for !slices.Equal(run(a), prog[i:]) {
			a++
		}
	}
	fmt.Println(a)
}

func main() {
	part1()
	part2()
}
