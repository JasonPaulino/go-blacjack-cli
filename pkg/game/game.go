package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Card struct {
	Suit   string
	Value  string
	Points int
}

type Deck []Card

type Player struct {
	Name  string
	Hand  []Card
	Score int
}

var (
	green   = color.New(color.FgGreen).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func showLogo() {
	logo := `
	██████╗ ██╗      █████╗  ██████╗██╗  ██╗     ██╗ █████╗  ██████╗██╗  ██╗
	██╔══██╗██║     ██╔══██╗██╔════╝██║ ██╔╝     ██║██╔══██╗██╔════╝██║ ██╔╝
	██████╔╝██║     ███████║██║     █████╔╝      ██║███████║██║     █████╔╝ 
	██╔══██╗██║     ██╔══██║██║     ██╔═██╗ ██   ██║██╔══██║██║     ██╔═██╗ 
	██████╔╝███████╗██║  ██║╚██████╗██║  ██╗╚█████╔╝██║  ██║╚██████╗██║  ██╗
	╚═════╝ ╚══════╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝ ╚════╝ ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝
	`
	fmt.Println(yellow(logo))
}

func animateDealing() {
	frames := []string{"🂠 ", " 🂠", "  🂠", "   🂠", "    🂠"}
	for i := 0; i < 5; i++ {
		for _, frame := range frames {
			fmt.Printf("\r%s", frame)
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println()
}

func createDeck() Deck {
	suits := []string{"♠", "♣", "♥", "♦"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	points := []int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	
	var deck Deck
	for _, suit := range suits {
		for i, value := range values {
			deck = append(deck, Card{suit, value, points[i]})
		}
	}
	return deck
}

func (d Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func calculateScore(hand []Card) int {
	score := 0
	aces := 0
	
	for _, card := range hand {
		if card.Value == "A" {
			aces++
		}
		score += card.Points
	}
	
	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}
	
	return score
}

func displayHand(hand []Card, hideFirst bool) {
	if hideFirst {
		fmt.Print(blue("🂠 "))
		for i := 1; i < len(hand); i++ {
			printCard(hand[i])
		}
		fmt.Println()
		return
	}
	
	for _, card := range hand {
		printCard(card)
	}
	fmt.Println()
}

func printCard(card Card) {
	cardStr := fmt.Sprintf("%s%s ", card.Value, card.Suit)
	if card.Suit == "♥" || card.Suit == "♦" {
		fmt.Print(red(cardStr))
	} else {
		fmt.Print(cyan(cardStr))
	}
}

func showResult(message string, isWin bool) {
	if isWin {
		fmt.Println(green("\n🎉 " + message + " 🎉"))
	} else {
		fmt.Println(red("\n💔 " + message + " 💔"))
	}
}

// StartGame starts a new blackjack game session
func StartGame() {
	clearScreen()
	showLogo()
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(magenta("\nEnter your name: "))
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSpace(playerName)
	
	for {
		clearScreen()
		showLogo()
		deck := createDeck()
		deck.shuffle()
		
		player := Player{Name: playerName, Hand: []Card{}, Score: 0}
		dealer := Player{Name: "Dealer", Hand: []Card{}, Score: 0}
		
		// Initial deal
		fmt.Println(yellow("\nDealing cards..."))
		animateDealing()
		
		player.Hand = append(player.Hand, deck[0], deck[1])
		dealer.Hand = append(dealer.Hand, deck[2], deck[3])
		deck = deck[4:]
		
		fmt.Printf("\n%s's hand: ", green(player.Name))
		displayHand(player.Hand, false)
		fmt.Printf("%s's hand: ", red(dealer.Name))
		displayHand(dealer.Hand, true)
		
		// Player's turn
		for {
			player.Score = calculateScore(player.Hand)
			if player.Score > 21 {
				break
			}
			
			fmt.Printf("\n%s's score: %s\n", green(player.Name), yellow(fmt.Sprintf("%d", player.Score)))
			fmt.Print(cyan("Do you want to (H)it or (S)tand? "))
			choice, _ := reader.ReadString('\n')
			choice = strings.ToUpper(strings.TrimSpace(choice))
			
			if choice == "H" {
				fmt.Println(yellow("\nDealing card..."))
				animateDealing()
				player.Hand = append(player.Hand, deck[0])
				deck = deck[1:]
				fmt.Printf("\n%s's hand: ", green(player.Name))
				displayHand(player.Hand, false)
			} else if choice == "S" {
				break
			}
		}
		
		// Dealer's turn
		fmt.Printf("\n%s's full hand: ", red(dealer.Name))
		displayHand(dealer.Hand, false)
		
		dealer.Score = calculateScore(dealer.Hand)
		for dealer.Score < 17 {
			fmt.Println(yellow("\nDealer hits..."))
			animateDealing()
			dealer.Hand = append(dealer.Hand, deck[0])
			deck = deck[1:]
			dealer.Score = calculateScore(dealer.Hand)
			fmt.Printf("%s hits: ", red(dealer.Name))
			displayHand(dealer.Hand, false)
		}
		
		// Determine winner
		player.Score = calculateScore(player.Hand)
		dealer.Score = calculateScore(dealer.Hand)
		
		fmt.Printf("\n%s Final scores %s\n", yellow("==="), yellow("==="))
		fmt.Printf("%s: %s\n", green(player.Name), yellow(fmt.Sprintf("%d", player.Score)))
		fmt.Printf("%s: %s\n", red(dealer.Name), yellow(fmt.Sprintf("%d", dealer.Score)))
		
		if player.Score > 21 {
			showResult("Bust! Dealer wins!", false)
		} else if dealer.Score > 21 {
			showResult(player.Name+" wins!", true)
		} else if player.Score > dealer.Score {
			showResult(player.Name+" wins!", true)
		} else if dealer.Score > player.Score {
			showResult("Dealer wins!", false)
		} else {
			fmt.Println(yellow("\n🤝 It's a tie! 🤝"))
		}
		
		fmt.Print(magenta("\nPlay again? (Y/N): "))
		again, _ := reader.ReadString('\n')
		again = strings.ToUpper(strings.TrimSpace(again))
		if again != "Y" {
			break
		}
	}
	
	fmt.Println(green("\nThanks for playing! 👋"))
}
