package main

import "fmt"

func main(){
	fmt.Println(1+1)	//tambah
	fmt.Println(5-2)	//kurang
	fmt.Println(10*2)	//kali
	fmt.Println(20/5)	//bagi
	fmt.Println(10%7)	//modulo

	//operasi logika AND (&&)
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	//operasi logika OR (||)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	//operasi logikan NOT (!)
	fmt.Println(!true)
	fmt.Println(!false)
}