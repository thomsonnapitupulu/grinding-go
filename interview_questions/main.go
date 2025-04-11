package main

import "fmt"

func f1() {
	defer fmt.Println("f1-begin")
	f2()
	defer fmt.Println("f1-end")
}
func f2() {
	defer fmt.Println("f2-begin")
	f3()
	defer fmt.Println("f2-end")
}
func f3() {
	defer fmt.Println("f3-begin")
	panic(0)
	defer fmt.Println("f3-end")
}

func main() {
	f1()
}

// // answer no. 1
// package main

// import "fmt"

// const N = 3

// func main() {
// 	m := make(map[int]*int)
// 	for i := 0; i < N; i++ {
// 		m[i] = &i
// 	}
// 	for _, v := range m {
// 		fmt.Println(*v)
// 	}
// }
