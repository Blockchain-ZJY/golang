package main

import "fmt"

const MOD = 1000000007

// 3623. 统计梯形的数目 I
func countTrapezoids(points [][]int) int {
	maps := make(map[int]int)
	for _, v := range points {
		maps[v[1]]++
	}
	pos := []int{}
	for _, v := range maps {
		pos = append(pos, (v%MOD*(v-1)%MOD/2)%MOD)
	}
	res := 0
	for i := 0; i < len(pos)-1; i++ {
		for j := i + 1; j < len(pos); j++ {
			res += (pos[i] * pos[j]) % MOD
		}
	}
	return res % MOD
}

// 3623. 统计梯形的数目 I
func countTrapezoids1(points [][]int) int {
	maps := make(map[int]int)
	for _, v := range points {
		maps[v[1]]++
	}
	pos := []int{}
	for _, v := range maps {
		pos = append(pos, (v*(v-1)/2)%MOD)
	}
	totaledges := 0
	for _, v := range pos {
		totaledges += v
	}
	res := 0
	for _, cur := range pos {
		res += (cur * (totaledges - cur)) % MOD
	}
	res = res * ((MOD + 1) / 2) % MOD
	return res
}

func main() {
	fmt.Println(countTrapezoids1([][]int{{0, 0}, {1, 0}, {0, 1}, {2, 2}}))
	fmt.Println(2998221387 % MOD)
}
