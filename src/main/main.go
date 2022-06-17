package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"strconv"
)

type GameScore struct{
	player_id string
	score int64
}

func main(){
	fmt.Println(os.Args[1])
	CSV := getCsv(os.Args[1])
	fmt.Println(CSV)
}

func getCsv(filename string)[]GameScore{
	file,err := os.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	var gameScores []GameScore

	rows,err := r.ReadAll()
	if err != nil{
		log.Fatal(err)
	}


	for i,row := range rows{
		if i == 0{
			continue
		}
		var t GameScore
		if i == 1{
			fmt.Println(row[0])
		}
		t.player_id = row[1]
		temp,_ := strconv.Atoi(row[2])
		t.score = int64(temp)
		gameScores = append(gameScores,t)
	}
	return gameScores
}