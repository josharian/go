package main

import (
	"fmt"
)

const loadFactor = 6.5

var lowbuckets = [...]int{0, 8, 13, 26, 52, 104, 208, 416, 832, 1664, 3328, 6656, 13312, 26624, 53248, 106496}

func main() {
	fmt.Println("START")
	m := make(map[int]string)
	for i := 0; i < 100000; i++ {
		m[i] = "i"
	}
}

/*

func main5() {
	var tots [16]int
	iters := 0
	for {
		for b, sz := range lowbuckets {
			m := make(map[int]struct{}, sz)
			for j := 0; j < sz; j++ {
				m[j] = struct{}{}
			}
			novf := runtime.MapNOverflow(*(*unsafe.Pointer)(unsafe.Pointer(&m)))
			tots[b] += int(novf)
		}
		iters++
		fmt.Println("ITERS", iters)
		for i := range tots {
			fmt.Printf("B=%d\tsz=%d\t\tnovf=%0.1f\n", i, lowbuckets[i], float64(tots[i])/float64(iters))
		}
		fmt.Println()
	}

	// ITERS 3563
	// B=0	sz=0		novf=0.0
	// B=1	sz=8		novf=0.0
	// B=2	sz=13		novf=0.0
	// B=3	sz=26		novf=0.0
	// B=4	sz=52		novf=0.1
	// B=5	sz=104		novf=0.2
	// B=6	sz=208		novf=0.4
	// B=7	sz=416		novf=0.8
	// B=8	sz=832		novf=1.6
	// B=9	sz=1664		novf=3.2
	// B=10	sz=3328		novf=6.4
	// B=11	sz=6656		novf=12.8
	// B=12	sz=13312		novf=25.6
	// B=13	sz=26624		novf=51.6
	// B=14	sz=53248		novf=103.1
	// B=15	sz=106496		novf=206.2

}*/

// truncated
var novf = [...]int{
	0:  0,
	1:  0,
	2:  0,
	3:  0,
	4:  0,
	5:  0,
	6:  0,
	7:  0,
	8:  1,
	9:  3,
	10: 6,
	11: 12,
	12: 25,
	13: 51,
	14: 103,
	15: 206,
}

// var novf = [...]int{
// 	0:  0.0,
// 	1:  0.0,
// 	2:  0.0,
// 	3:  0.0,
// 	4:  0.1,
// 	5:  0.2,
// 	6:  0.4,
// 	7:  0.8,
// 	8:  1.6,
// 	9:  3.2,
// 	10: 6.4,
// 	11: 12.8,
// 	12: 25.6,
// 	13: 51.6,
// 	14: 103.1,
// 	15: 206.2,
// }

func main3() {
	var prev uint8
	for n := 0; n < 20000000; n++ {
		x := nbuckets(int64(n))
		if x != prev {
			fmt.Println("size=", n, "buckets=", x)
			prev = x
		}
	}
	// size= 8 buckets= 1
	// size= 13 buckets= 2
	// size= 26 buckets= 3
	// size= 52 buckets= 4
	// size= 104 buckets= 5
	// size= 208 buckets= 6
	// size= 416 buckets= 7
	// size= 832 buckets= 8
	// size= 1664 buckets= 9
	// size= 3328 buckets= 10
	// size= 6656 buckets= 11
	// size= 13312 buckets= 12
	// size= 26624 buckets= 13
	// size= 53248 buckets= 14
	// size= 106496 buckets= 15
	// size= 212992 buckets= 16
	// size= 425984 buckets= 17
	// size= 851968 buckets= 18
	// size= 1703936 buckets= 19
	// size= 3407872 buckets= 20
	// size= 6815744 buckets= 21
	// size= 13631488 buckets= 22
}

func main2() {
	for i := 0; i < 1000; i++ {
		m := make(map[int]struct{})
		for j := 0; j < 1000000; j++ {
			m[j] = struct{}{}
		}
	}
	// fmt.Println("size\tratio\toverflow\tB\testimate\tacc\test2\tacc2\test3\tacc3")
	// for sz := 100; sz <= 10000; sz += 100 {
	// 	f := testing.AllocsPerRun(100, func() {
	// 		m := make(map[int]struct{}, sz)
	// 		for i := 0; i < sz; i++ {
	// 			m[i] = struct{}{}
	// 		}
	// 	})

	// 	B := nbuckets(int64(sz))

	// 	ratio := float32(sz) / (loadFactor * float32((uintptr(1) << B)))

	// 	c1 := coeffs1[B]
	// 	c2 := coeffs2[B]
	// 	c3 := coeffs3[B]
	// 	est1 := c1.estimate(sz)
	// 	est2 := c2.estimate(sz)
	// 	est3 := c3.estimate(sz)

	// 	fmt.Printf("%d\t%f\t%d\t%d\t%d\t%0.2f\t%d\t%0.2f\t%d\t%0.2f\n", sz, ratio, int(f), B, est1, float64(est1)/f, est2, float64(est2)/f, est3, float64(est3)/f)
	// }
}

// const (
// 	c1_0 = 2.2102038681951472e-006
// 	c1_1 = -0.11474100740321222
// 	c1_2 = 1566.3789153164603
// )

// var coeffs1 = [...]coeff{
// 	0:  [3]float64{c1_0 * 8192, c1_1, c1_2 / 8192},
// 	1:  [3]float64{c1_0 * 4096, c1_1, c1_2 / 4096},
// 	2:  [3]float64{c1_0 * 2048, c1_1, c1_2 / 2048},
// 	3:  [3]float64{c1_0 * 1024, c1_1, c1_2 / 1024},
// 	4:  [3]float64{c1_0 * 512, c1_1, c1_2 / 512},
// 	5:  [3]float64{c1_0 * 256, c1_1, c1_2 / 256},
// 	6:  [3]float64{c1_0 * 128, c1_1, c1_2 / 128},
// 	7:  [3]float64{c1_0 * 64, c1_1, c1_2 / 64},
// 	8:  [3]float64{c1_0 * 32, c1_1, c1_2 / 32},
// 	9:  [3]float64{c1_0 * 16, c1_1, c1_2 / 16},
// 	10: [3]float64{c1_0 * 8, c1_1, c1_2 / 8},
// 	11: [3]float64{c1_0 * 4, c1_1, c1_2 / 4},
// 	12: [3]float64{c1_0 * 2, c1_1, c1_2 / 2},
// 	13: [3]float64{c1_0, c1_1, c1_2},
// 	14: [3]float64{c1_0 / 2, c1_1, c1_2 * 2},
// }

// const (
// 	c2_0 = 1.0899364166233374e-006
// 	c2_1 = -0.11258320807498451
// 	c2_2 = 3047.4987992501583
// )

// var coeffs2 = [...]coeff{
// 	[3]float64{c2_0 * 16384, c2_1, c2_2 / 16384},
// 	[3]float64{c2_0 * 8192, c2_1, c2_2 / 8192},
// 	[3]float64{c2_0 * 4096, c2_1, c2_2 / 4096},
// 	[3]float64{c2_0 * 2048, c2_1, c2_2 / 2048},
// 	[3]float64{c2_0 * 1024, c2_1, c2_2 / 1024},
// 	[3]float64{c2_0 * 512, c2_1, c2_2 / 512},
// 	[3]float64{c2_0 * 256, c2_1, c2_2 / 256},
// 	[3]float64{c2_0 * 128, c2_1, c2_2 / 128},
// 	[3]float64{c2_0 * 64, c2_1, c2_2 / 64},
// 	[3]float64{c2_0 * 32, c2_1, c2_2 / 32},
// 	[3]float64{c2_0 * 16, c2_1, c2_2 / 16},
// 	[3]float64{c2_0 * 8, c2_1, c2_2 / 8},
// 	[3]float64{c2_0 * 4, c2_1, c2_2 / 4},
// 	[3]float64{c2_0 * 2, c2_1, c2_2 / 2},
// 	[3]float64{c2_0, c2_1, c2_2},
// }

// const (
// 	c3_0 = 2.2e-06
// 	c3_1 = -0.1135
// 	c3_2 = 1545
// )

// var coeffs3 = [...]coeff{
// 	0:  [3]float64{c3_0 * 8192, c3_1, c3_2 / 8192},
// 	1:  [3]float64{c3_0 * 4096, c3_1, c3_2 / 4096},
// 	2:  [3]float64{c3_0 * 2048, c3_1, c3_2 / 2048},
// 	3:  [3]float64{c3_0 * 1024, c3_1, c3_2 / 1024},
// 	4:  [3]float64{c3_0 * 512, c3_1, c3_2 / 512},
// 	5:  [3]float64{c3_0 * 256, c3_1, c3_2 / 256},
// 	6:  [3]float64{c3_0 * 128, c3_1, c3_2 / 128},
// 	7:  [3]float64{c3_0 * 64, c3_1, c3_2 / 64},
// 	8:  [3]float64{c3_0 * 32, c3_1, c3_2 / 32},
// 	9:  [3]float64{c3_0 * 16, c3_1, c3_2 / 16},
// 	10: [3]float64{c3_0 * 8, c3_1, c3_2 / 8},
// 	11: [3]float64{c3_0 * 4, c3_1, c3_2 / 4},
// 	12: [3]float64{c3_0 * 2, c3_1, c3_2 / 2},
// 	13: [3]float64{c3_0, c3_1, c3_2},
// 	14: [3]float64{c3_0 / 2, c3_1, c3_2 * 2},
// }

func nbuckets(sz int64) uint8 {
	B := uint8(0)
	for ; overLoadFactor(sz, B); B++ {
	}
	return B
}

func overLoadFactor(count int64, B uint8) bool {
	// TODO: rewrite to use integer math and comparison?
	const bucketCnt = 8
	return count >= bucketCnt && float32(count) >= loadFactor*float32((uintptr(1)<<B))
}

// func lowbucketestimate(n int) int {
// 	return int(5.61157 + 313.3373838234561*(1-math.Exp(-0.00000742297*float64(n))))
// }

// func highbucketestimate(n int) int {
// 	return int(4.932299 + 78573.64843833569*(1-math.Exp(-4.150828e-7*float64(n))))
// }

// type coeff [3]float64

// func (c coeff) estimate(n int) int {
// 	f := float64(n)
// 	return int(c[0]*f*f + c[1]*f + c[2])
// }
