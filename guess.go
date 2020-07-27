package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I got a random number between 1 ~ 100")
	fmt.Println("You gotta guess it. You have 10 chances")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10 - guesses, "chances left")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("HAHA your guess is too LOW")
		} else if guess > target {
			fmt.Println("HAHA your guess it too HIGH")
		} else {
			success = true
			fmt.Println("Good job you guessed it correct")
			break
		}
	}
	if !success {
		fmt.Println("You didnt guess my number right. I will EAT YOU!! number was: ", target)
	}
}