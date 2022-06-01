package main

import (
	"fmt"
	"os"
	"strconv"
)

func calLunarYear(year int) {
	can := [10]string{"Canh","Tan","Nham","Quy","Giap","At","Binh","Dinh","Mau","Ky"}
	chi := [12]string{"Than","Dau","Tuat","Hoi","Ty'","Suu","Dan","Meo","Thin","Ty.","Ngo","Mui"}

	thu_tu_can := year % 10
	thu_tu_chi := year % 12

	fmt.Println("Năm", year, "là năm", can[thu_tu_can], chi[thu_tu_chi])
}


func main() {
	year := os.Args[1]
	y, _ := strconv.Atoi(year)
	calLunarYear(y)
}