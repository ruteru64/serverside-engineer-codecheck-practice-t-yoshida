package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"strconv"
	"sort"
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


// データを取得しそのデータからランキングを取得する
func main(){
	output := getPlayerScore(getCsv(os.Args[1]))

	sort.SliceStable(output, func(i, j int) bool { return output[i].player_id < output[j].player_id })
	sort.SliceStable(output, func(i, j int) bool { return output[i].score_average > output[j].score_average })
	rank := 1
	fmt.Println("rank,player_id,mean_score")
	for i:=0;rank<=10;i++{
		if i >= len(output){
			return
		}
		fmt.Printf("%d,%s,%d\n",rank,output[i].player_id,output[i].score_average)
		if i+1 >= len(output){
			return
		}
		if output[i].score_average == output[i+1].score_average{
		}else{
			rank = i+2
		}
	}
}


// ログからそれぞれのプレイヤーのスコアを取得する
// @param gameScores 取得したゲームのログ
// @return それぞれのプレイヤーのスコア
func getPlayerScore(gameScores []GameScore)[]PrayerScore{
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
	return output
}

// Csvを取得しフォーマットする
// @param filename 取得するファイル名
// @return 取得したCSVをフォーマットしたもの
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
		t.player_id = row[1]
		temp,_ := strconv.Atoi(row[2])
		t.score = int64(temp)
		gameScores = append(gameScores,t)
	}
	return gameScores
}


// 配列の中から指定した文字列を探す
// @pram arr 探す配列
// @pram str 探す文字列
// @return 見つかった時配列のindex、見つからなかった時-1
func arrayContains(arr []PrayerScore, str string) int{
	for i, v := range arr{
	  if v.player_id == str{
		return i
	  }
	}
	return -1
}