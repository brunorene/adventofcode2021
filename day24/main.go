package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func x(index int, w, z int64) int64 {
	val := [14]int64{11, 13, 12, 15, 10, -1, 14, -8, -7, -8, 11, -2, -2, -13}
	if (z%26)+val[index] != w {
		return 1
	}
	return 0
}

func z0(index int, z int64) int64 {
	if index < 5 || index == 6 || index == 10 {
		return z
	}
	return z / 26
}

func y0(index int, w, z int64) int64 {
	return 25*x(index, w, z) + 1
}

func z1(index int, w, z int64) int64 {
	return z0(index, z) * y0(index, w, z)
}

func y1(index int, w, z int64) int64 {
	val := [14]int64{5, 5, 1, 15, 2, 2, 5, 8, 14, 12, 7, 14, 13, 6}
	return (w + val[index]) * x(index, w, z)
}

func zFinal(index int, w, z int64) int64 {
	return z1(index, w, z) + y1(index, w, z)
}

func calcZ(cache map[[3]int64]int64, digits ...int64) (z int64) {
	var exists bool
	for idx, n := range digits {
		z, exists = cache[[3]int64{int64(idx), n, z}]
		if !exists {
			oldZ := z
			z = zFinal(idx, n, z)
			cache[[3]int64{int64(idx), n, oldZ}] = z
		}
	}

	return
}

func minValid() {
	cache := make(map[[3]int64]int64)
	for n0 := int64(1); n0 < 10; n0++ {
		for n1 := int64(1); n1 < 10; n1++ {
			for n2 := int64(1); n2 < 10; n2++ {
				for n3 := int64(1); n3 < 10; n3++ {
					for n4 := int64(1); n4 < 10; n4++ {
						for n5 := int64(1); n5 < 10; n5++ {
							for n6 := int64(1); n6 < 10; n6++ {
								for n7 := int64(1); n7 < 10; n7++ {
									for n8 := int64(1); n8 < 10; n8++ {
										for n9 := int64(1); n9 < 10; n9++ {
											for n10 := int64(1); n10 < 10; n10++ {
												for n11 := int64(1); n11 < 10; n11++ {
													for n12 := int64(1); n12 < 10; n12++ {
														for n13 := int64(1); n13 < 10; n13++ {
															// n0 := int64(1)
															// n1 := int64(1)
															// n2 := int64(9)
															// n3 := int64(1)
															// n4 := int64(8)
															// n5 := int64(9)
															// n6 := int64(9)
															// n7 := int64(6)
															// n8 := int64(9)
															// n9 := int64(2)
															// n10 := n10and11[0]
															// n11 := n10and11[1]
															// n12 := int64(1)
															// n13 := int64(1)
															z := calcZ(cache, n0, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13)
															if z < 1000000000000000000 {
																fmt.Printf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d - %d\n",
																	n0, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13, z)
																// bufio.NewReader(os.Stdin).ReadBytes('\n')
															}
															if z == 0 {
																fmt.Printf("*** %d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",
																	n0, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

}

func maxValid() {
	cache := make(map[[3]int64]int64)
	for n0 := int64(9); n0 >= 1; n0-- {
		for n1 := int64(9); n1 >= 1; n1-- {
			for n2 := int64(9); n2 >= 1; n2-- {
				// for n3 := int64(9); n3 >= 1; n3-- {
				// 	for n4 := int64(9); n4 >= 1; n4-- {
				// 		for n5 := int64(9); n5 >= 1; n5-- {
				// 			for n6 := int64(9); n6 >= 1; n6-- {
				// 				for n7 := int64(9); n7 >= 1; n7-- {
				// 					for n8 := int64(9); n8 >= 1; n8-- {
				// 						for n9 := int64(9); n9 >= 1; n9-- {
				// 							for n10 := int64(9); n10 >= 1; n10-- {
				// 								for n11 := int64(9); n11 >= 1; n11-- {
				for n12 := int64(9); n12 >= 1; n12-- {
					for n13 := int64(9); n13 >= 1; n13-- {
						// 119 189969249 99
						// n0 := int64(1)
						// n1 := int64(1)
						// n2 := int64(9)
						n3 := int64(1)
						n4 := int64(8)
						n5 := int64(9)
						n6 := int64(9)
						n7 := int64(6)
						n8 := int64(9)
						n9 := int64(2)
						n10 := int64(4)
						n11 := int64(9)
						// n12 := int64(9)
						// n13 := int64(9)
						z := calcZ(cache, n0, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13)
						fmt.Printf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d - %d\n",
							n0, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13, z)
						// bufio.NewReader(os.Stdin).ReadBytes('\n')
						if z == 0 {
							fmt.Printf("*** %d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",
								n0, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13)
							return
						}
						// }
						// }
						// }
					}
				}
			}
		}
	}
	// }
	// 	}
	// }
	// }
	// }
	// }

}

func part1() {
	maxValid()
}

func part2() {
	// minValid()
}

func main() {
	part1()
	part2()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(filename string) (lines []string) {
	_, path, _, _ := runtime.Caller(0)
	dir := strings.ReplaceAll(path, "main.go", "")

	file, err := os.Open(dir + filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())

	return
}
