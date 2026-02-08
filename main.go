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

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

const (
	selectionsPerChance = 5
	maxChances          = 5
)

type Player struct {
	Name   string
	Games  int
	Wins   int
	Scores int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	displayBanner()

	fmt.Print("Enter your heroic name: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	if name == "" {
		name = "Adventurer"
	}

	player := Player{Name: name}

	for {
		fmt.Println("\nMENU")
		fmt.Println("1) Play round")
		fmt.Println("2) View scores")
		fmt.Println("3) Help / Instructions")
		fmt.Println("4) Free try")
		fmt.Println("5) Exit")
		fmt.Print("Enter choice (1-5): ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			playRound(&player, scanner)
		case "2":
			showScores(&player)
		case "3":
			showHelp()
		case "4":
			fmt.Println("\nFREE TRY ROUND")
			playRound(&player, scanner)
		case "5":
			fmt.Println(Cyan + "Farewell, brave soul!" + Reset)
			return
		default:
			fmt.Println(Red + "Invalid choice. Enter 1-5." + Reset)
		}
	}
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
	fmt.Println(Cyan + "WELCOME TO LETTER QUEST" + Reset)
}

func showHelp() {
	fmt.Println("\nHELP / INSTRUCTIONS")
	fmt.Println("Match alphanumeric values and win a prize!")
	fmt.Println("\nAlphabet Mapping:")
	fmt.Println("A -> 1, B -> 2, C -> 3, ...")
	fmt.Println("X -> 24, Y -> 25, Z -> 26")
	fmt.Println("\nHow to Play:")
	fmt.Println("• You get 5 chances per round.")
	fmt.Println("• Each chance allows 5 number selections.")
	fmt.Println("• If a number matches its letter value, you score.")
	fmt.Println("• 3 or more total matches = WIN (score multiplied by 5).")
	fmt.Println("• Multiples of 5 or number 26 give royal blessings.")
}

func showScores(player *Player) {
	fmt.Printf("\nSTATS FOR %s\n", player.Name)
	fmt.Printf("Games played: %d\n", player.Games)
	fmt.Printf("Wins: %d\n", player.Wins)
	fmt.Printf("Total score: %d\n", player.Scores)
}

func playRound(player *Player, scanner *bufio.Scanner) {
	player.Games++

	letters, numbers := generateSets()
	shuffle(letters)
	shuffle(numbers)

	totalCorrect := 0

	for chance := 1; chance <= maxChances; chance++ {
		fmt.Printf("\nChance %d of %d — select 5 numbers: ", chance, maxChances)
		scanner.Scan()
		input := strings.Fields(scanner.Text())

		if len(input) != selectionsPerChance {
			fmt.Println(Red + "You must select exactly 5 numbers." + Reset)
			chance--
			continue
		}

		selections := make([]int, 0)
		for _, v := range input {
			n, err := strconv.Atoi(v)
			if err != nil || n < 1 || n > 26 {
				fmt.Println(Red + "Invalid number entered." + Reset)
				selections = nil
				break
			}
			selections = append(selections, n)
		}

		if selections == nil {
			chance--
			continue
		}

		correct := revealLetters(letters, numbers, selections)
		totalCorrect += correct
	}

	score := totalCorrect * 10

	if totalCorrect >= 3 {
		player.Wins++
		score *= 5
		player.Scores += score
		fmt.Println(Green + "\nROYAL BLESSINGS!" + Reset)
		displayAfricanMask()
		deities := []string{"Ogun", "Shango", "Yemoja", "Orunmila", "Obatala"}
		fmt.Printf(Green+"%s blesses you, %s! Keep playing!\n"+Reset,
			deities[rand.Intn(len(deities))], player.Name)
	} else {
		player.Scores += score
		fmt.Println(Red + "\nYou lose, try again!" + Reset)
	}

	fmt.Printf("\nTotal matches: %d\n", totalCorrect)
	fmt.Printf("Round score: %d | Total score: %d\n", score, player.Scores)
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

func shuffle[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func revealLetters(letters []rune, numbers []int, picks []int) int {
	correct := 0
	fmt.Println("\nRevealing letters:")
	for i, l := range letters {
		n := numbers[i]
		for _, p := range picks {
			if p == n {
				if n == int(l-'A')+1 {
					fmt.Printf(Green+"%c -> %d\n"+Reset, l, n)
					correct++
				} else {
					fmt.Printf(Yellow+"%c -> %d\n"+Reset, l, n)
				}
				break
			}
		}
	}
	return correct
}

func displayAfricanMask() {
	fmt.Println(Yellow + `
   .-''''-.
  /  .--.  \
 |  |    |  |
 |  |.-""-. |
 |||  ____  |
 ||| |____| |
  \\\        /
   '-......-'
` + Reset)
}
