package main

import (
	"fmt"
	"strings"
)

func main(){
	var arr []string = []string {"TeamA-win", "TeamB-draw", "TeamC-goal", "TeamD-loss", "TeamE-draw"}
	calculateTeamPoint(arr)

}

func calculateTeamPoint(array []string){
	for _, team := range array{
		t := strings.Split(team, "-")
		if t [1] =="win"{
			fmt.Printf("Point for %s is %d\n",t[0], 3)
		} else if t [1] == "draw" {
			fmt.Printf("Point for %s is %d\n",t [0], 1)
		}else if t [1] == "loss"{
			fmt.Printf("Point for %s is %d\n", t[0], 0)
		}else {
			fmt.Println("Invalid score")
		}
		
		
		
	}
}

