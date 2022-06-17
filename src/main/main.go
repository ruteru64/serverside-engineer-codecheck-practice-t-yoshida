package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
)

func main(){
	fmt.Println(os.Args[1])
	CSV := getCsv(os.Args[1])
	fmt.Println(CSV)
}

func getCsv(filename string)[][]string{
	file,err := os.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows,err := r.ReadAll()

	if err != nil{
		log.Fatal(err)
	}

	return rows
}