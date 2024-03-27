package main

import "fmt"

func main()  {
	var name, decision string
	var score, ans int
	fmt.Println(`Welcome to my math game!`)
	fmt.Println("First things first tell me your name:")
	fmt.Scanln(&name)
	fmt.Print("Okay ", name, ", shall we begin? (y/n) \n")
	fmt.Scanln(&decision)
	Checkdecision(decision)
	fmt.Print("First question: What's 5 + 8? ")
	fmt.Scan(&ans);
	if ans == 13  {
		fmt.Println("That's correct!")
		score++
	} else {
		fmt.Println("I'm sure you will answer correctly next time")
	}
	fmt.Print("Next one: What's 89 - 23? ")
	fmt.Scan(&ans)
	if ans == 66 {
		fmt.Println("Great!")
		score++
	} else {
		fmt.Println("Oops, not this time.")
	}
	fmt.Print("Last question: What's 9 + 10? ")
	fmt.Scan(&ans)
	if ans == 19 {
		fmt.Println("That's right!")
		score++
	} else if ans == 21 {
		fmt.Println("I also like this meme, you should have this point.")
		score++
	} else {
		fmt.Println("That's not quite it.")
	}
	fmt.Println("Okay, that's it for the game. To win you need at least 2 points. Your score is:", score)
	if score > 1 {
		fmt.Println("Congratulations! You've won!")
	} else {
		fmt.Println("You didn't quite make it, but you can try again!")
	}
}

// function that checks if user wants to begin the game
// no matter what they choose they are forced to begin ;)
func Checkdecision(decision string)  {
	if decision == "n" {
		fmt.Println("Nonetheless, we will begin the game.")
	} else if decision == "y" {
		fmt.Println("Very well.")
	} else {
		fmt.Println("I don't understand if it isn't y or n. Let's forget about it and begin.")
	}
}