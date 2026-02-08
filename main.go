package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ANSI Colors
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

const selectionCount = 5

type Player struct {
	Name   string
	Wins   int
	Games  int
	Scores int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	displayBanner()

	fmt.Print("Enter your heroic name: ")
	scanner.Scan()
	name := scanner.Text()
	if name == "" {
		name = "Adventurer"
	}

	player := Player{Name: name}

	for {
		fmt.Println("\n🎯 Select 5 numbers (1-26). Match the hidden letters to win treasures!")
		letters, numbers := generateSets()
		shuffle(letters)
		shuffle(numbers)
		masks := generateMasks(len(letters))

		playRound(&player, scanner, letters, numbers, masks)

		fmt.Print("\nDo you wish to venture again? (yes/no): ")
		scanner.Scan()
		if strings.ToLower(strings.TrimSpace(scanner.Text())) != "yes" {
			fmt.Println(Cyan + "Farewell, brave soul!" + Reset)
			break
		}
	}

	fmt.Printf("\nGames played: %d | Wins: %d | Total Score: %d\n", player.Games, player.Wins, player.Scores)
}

func displayBanner() {
	fmt.Println(Cyan + `
████████╗ █████╗ █████╗ █████╗
╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗
   ██║   ███████║███████║███████║
   ██║   ██╔══██║╚════██║██╔══██║
   ██║   ██║  ██║███████║██║  ██║
   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
` + Reset)
	fmt.Println(Cyan + "🌟 Welcome to LETTER QUEST 🌟" + Reset)
}

func generateSets() ([]rune, []int) {
	letters := make([]rune, 26)
	numbers := make([]int, 26)
	for i := 0; i < 26; i++ {
		letters[i] = rune('A' + i)
		numbers[i] = i + 1
	}
	return letters, numbers
}

func shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func generateMasks(n int) []string {
	symbols := []string{"[?]", "[#]", "[*]", "[!]", "[%]", "[@]", "[&]", "[~]"}
	masks := make([]string, n)
	for i := 0; i < n; i++ {
		masks[i] = symbols[rand.Intn(len(symbols))]
	}
	return masks
}

func playRound(player *Player, scanner *bufio.Scanner, letters []rune, numbers []int, masks []string) {
	player.Games++
	fmt.Print("Your selections: ")
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	if len(input) != selectionCount {
		fmt.Println(Red + "You must pick exactly 5 numbers!" + Reset)
		return
	}

	selections := make([]int, selectionCount)
	for i, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil || num < 1 || num > 26 {
			fmt.Println(Red+"Invalid number:", v+Reset)
			return
		}
		selections[i] = num
	}

	correct := 0
	fmt.Println("\n🔍 Revealing hidden letters:")
	for i, letter := range letters {
		num := numbers[i]
		selected := false
		for _, s := range selections {
			if s == num {
				selected = true
				break
			}
		}
		display := fmt.Sprintf("%c->%d", letter, num)
		if selected && num == int(letter-'A')+1 {
			correct++
			fmt.Println(Green + display + Reset)
		} else if selected {
			fmt.Println(Yellow + display + Reset)
		} else {
			fmt.Println(display)
		}
	}

	// Base score: 10 points per correct number
	roundScore := correct * 10

	// If player wins (matches 3 or more), show mask, congratulate, multiply score
	if correct >= 3 {
		player.Wins++
		roundScore *= 5
		fmt.Println(Green + "\n🎉 CONGRATULATIONS! You won this round!" + Reset)
		displayAfricanMask()
	}

	player.Scores += roundScore
	fmt.Printf("\nYou matched %d out of 5 numbers\n", correct)
	fmt.Printf("Round score: %d | Total score: %d\n", roundScore, player.Scores)

	// Bonus for special numbers
	for _, n := range selections {
		if n%5 == 0 || n == 26 {
			fmt.Println(Cyan + "✨ You received a royal blessing!" + Reset)
		}
	}
}

// African mask ASCII art
func displayAfricanMask() {
	mask := `
   .-''''-.
  /  .--.  \
 /  /    \  \
 |  |    |  |
 |  |.-""-. |
 ///`. + "`" + `::::.` + "`" + `\\\
||| ::/  \:: ;|
||| ::\__/:: ;|
 \\\ '::::' ///
  ` + "`" + `'-....-'`
	fmt.Println(Yellow + mask + Reset)
}