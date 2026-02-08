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

const maxChances = 5
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

	displayWelcomeBanner()

	fmt.Print("Enter your name, brave learner: ")
	scanner.Scan()
	username := scanner.Text()
	if username == "" {
		username = "Traveler"
	}

	player := Player{Name: username}

	for {
		fmt.Println("\nHow to play: Select 5 numbers corresponding to hidden letters (A->1 ... Z->26).")
		fmt.Println("3 correct → win, 4 correct → ancestral royalty, 5 correct → kingship.")
		fmt.Println("Special numbers 5,10,15,20,25,26 → royal blessing!")

		letters, numbers := generateSets()
		shuffle(letters)
		shuffle(numbers)
		masks := generateMasks(len(letters))

		playRound(&player, scanner, letters, numbers, masks)

		fmt.Println("\nDo you want to play again? (yes/no):")
		if !scanner.Scan() || strings.ToLower(strings.TrimSpace(scanner.Text())) != "yes" {
			fmt.Println(Cyan + "Thank you for playing! Return for revenge and glory!" + Reset)
			break
		}

		shuffle(letters)
		shuffle(numbers)
	}
	fmt.Printf("Games played: %d, Wins: %d, Total score: %d\n", player.Games, player.Wins, player.Scores)
}

// ---------------- BANNERS ----------------
func displayWelcomeBanner() {
	displayTata2D()
	fmt.Println(Cyan + "🎮 Welcome to INSIBIDICODE 🎮" + Reset)
}

func displayTata2D() {
	fmt.Println(`
TTTTT     A       TTTTTTT     A
  T     A A A        T      A A A
  T    A A A A       T     A A A A
  T   AAAAAAAAAA     T    AAAAAAAAAA
  T   A        A     T    A         A
`)
}

// ---------------- GAME LOGIC ----------------
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

func generateMasks(n int) []string {
	maskSymbols := []string{"[M]", "[X]", "[#]", "[*]", "[%]", "[&]", "[@]", "[!]", "[?]", "[~]"}
	masks := make([]string, n)
	for i := 0; i < n; i++ {
		masks[i] = maskSymbols[rand.Intn(len(maskSymbols))]
	}
	return masks
}

func playRound(player *Player, scanner *bufio.Scanner, letters []rune, numbers []int, masks []string) {
	player.Games++
	fmt.Println("\nSelect 5 numbers (1-26) separated by space:")
	fmt.Print("Your choices: ")
	if !scanner.Scan() {
		fmt.Println("\nGame aborted.")
		return
	}

	input := strings.Fields(scanner.Text())
	if len(input) != selectionCount {
		fmt.Println(Red + "You must select exactly 5 numbers!" + Reset)
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
	fmt.Println("\nHidden masks revealed:")
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

	score := correct * 10
	player.Scores += score

	fmt.Printf("\nYou matched %d out of 5\n", correct)

	// Rewards
	if correct == 3 {
		fmt.Println(Green + "You win! The ancestors smile upon you!" + Reset)
		player.Wins++
	} else if correct == 4 {
		fmt.Println(Green + "Ancestral royalty blesses you!" + Reset)
		fmt.Println(Yellow + "Wisdom: 'Knowledge is the bridge to power.'" + Reset)
		player.Wins++
	} else if correct == 5 {
		fmt.Println(Green + "KINGS AND QUEENS! You are the right hand of the king!" + Reset)
		player.Wins++
	}

	// Special numbers reward
	for _, n := range selections {
		if n == 5 || n == 10 || n == 15 || n == 20 || n == 25 || n == 26 {
			fmt.Println(Cyan + "Royal blessing from African heroes!" + Reset)
		}
	}
}
