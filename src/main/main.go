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

type PrayerScore struct{
	player_id string
	score_sum int64
	count int64
	score_average int64
}

func main(){
	fmt.Println(os.Args[1])
	gameScores := getCsv(os.Args[1])
	var output []PrayerScore
	for _,gameScore := range gameScores{
		index := arrayContains(output,gameScore.player_id)
		if index == -1{// 配列ないに同じ値がない時
			output = append(output,PrayerScore{gameScore.player_id,gameScore.score,1,0})
		}else{
			output[index].count++
			output[index].score_sum += gameScore.score
		}
	}

	for i,prayerScore := range output{
		output[i].score_average = (prayerScore.score_sum + prayerScore.count/2) / prayerScore.count
	}
	fmt.Println(output)
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
		if i == 0{// 一行目を握りつぶす
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

func arrayContains(arr []PrayerScore, str string) int{
	for i, v := range arr{
	  if v.player_id == str{
		return i
	  }
	}
	return -1
  }