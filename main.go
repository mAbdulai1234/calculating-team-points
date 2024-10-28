/*
Challenge 2: Calculate Team Points

Write a code that takes an array of strings representing the results of football matches and calculates the total points for a team.
Each string in the array will have the format
"TeamA-TeamB scoreA-scoreB". A team gets 3 points for a win, 1 point for a draw, and 0 points for a loss.
*/

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

