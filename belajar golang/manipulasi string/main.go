package main

import "fmt"

func main(){
	fmt.Println("belajar golang")
	fmt.Println("belajar" + " " + "go")
	fmt.Println(len("belajar"))//menganmbil panjang string
	fmt.Println("belajar golang"[0])//mengambil data karakter
	fmt.Println("belajar golang"[0:5]) //mengambil beberapa data karakter
	fmt.Println("belajar golang"[3:]) //mengambil awal data karakter ke akhir
	fmt.Println("belajar golang"[:10]) //mengambil data karakter  ke akhir
}
