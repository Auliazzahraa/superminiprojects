package main

import(
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func game(input int, num *int, score *int){
	var x int
	var answer int
	if input == 1{
		x = 10
	}else if input == 2{
		x = 5
	}else {
		x = 3
	}
	randNumber := rand.Intn(100)
	*num = randNumber + 1
	guess:=x
	for j:=1; j<=x;j++{
		fmt.Print("Enter your guess: ")
		fmt.Scan(&answer)
		if answer > *num {
			fmt.Println("Incorrect! The number is less than ", answer)
			fmt.Println()
		}else if answer < *num {
			fmt.Println("Incorrect! The number is greater than ", answer)
			fmt.Println()
		}else{
			fmt.Println("Congratulations! You guessed the correct number in ",j," attempts")
			*score = j
			break
		}
		
		if x > 3 && j%2 == 0{
			var hint string
			fmt.Print("Want a hint? (yes/no): ")
			fmt.Scan(&hint)
			if strings.ToLower(hint) == "yes"{
				rangeMin:= *num - guess
				rangeMax := *num + guess
				fmt.Printf("The answer is in between: [%v...%v]\n",rangeMin,rangeMax)
			}
		}
		
		guess--
		fmt.Printf("----------- %v guess left -----------\n",guess)
		
	}
}

func rounds(rounds int,input int, num int){
	var cont string
	var score, min int
	min = 15 //to set it max first
	for i:=1; i<=rounds;i++{
		fmt.Printf("===================== Round %v =====================\n",i)
		game(input, &num, &score)
		if min > score{
			min = score
		}
		fmt.Println("The answer: ", num)
		fmt.Println()
		if i!=rounds{
			fmt.Print("Do you still wanna play? (yes/no): ")
			fmt.Scan(&cont)
			if strings.ToLower(cont) == "no"{
				fmt.Print("Your best score = ",min)
				break
			}
		}
	}
	fmt.Println("Your best score = ",min)
}

func main(){
	var input, num, round int
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println()
	fmt.Print("How many rounds do you want?: ")
	fmt.Scan(&round)
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances, )")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&input)
	fmt.Println()
	if input == 1 || input == 2 || input == 3{
		startTime := time.Now()
		rounds(round, input, num)
		elapsedTime := time.Since(startTime)
		fmt.Printf("Time : %.2f seconds\n", elapsedTime.Seconds())
	}else{
		fmt.Print("Input invalid")
	}
}