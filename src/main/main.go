package main

import (
	"io"
	"time"
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"strconv"
)

const timeFormat = "2021/01/01 12:00"

type GameScore struct{
	create_timestamp time.Time
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

	// 1行目を飛ばす
	r.Read()


	for{
		row,err := r.Read()
		if err != io.EOF{
			break
		}
		var t GameScore
		t.create_timestamp = stoTime(row[0])
		t.player_id = row[1]
		temp,_ := strconv.Atoi(row[2])
		t.score = int64(temp)
		gameScores = append(gameScores,t)
	}
	return gameScores
}

func stoTime(s string)time.Time{
	t,_ := time.Parse(timeFormat,s)
	return t
}

func timetos(t time.Time)string{
	s := t.Format(timeFormat)
	return s
}