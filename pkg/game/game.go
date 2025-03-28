package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
		fmt.Print("🂠 ")
		for i := 1; i < len(hand); i++ {
			fmt.Printf("%s%s ", hand[i].Value, hand[i].Suit)
		}
		fmt.Println()
		return
	}
	
	for _, card := range hand {
		fmt.Printf("%s%s ", card.Value, card.Suit)
	}
	fmt.Println()
}

// StartGame starts a new blackjack game session
func StartGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Blackjack!")
	fmt.Print("Enter your name: ")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSpace(playerName)
	
	for {
		deck := createDeck()
		deck.shuffle()
		
		player := Player{Name: playerName, Hand: []Card{}, Score: 0}
		dealer := Player{Name: "Dealer", Hand: []Card{}, Score: 0}
		
		// Initial deal
		player.Hand = append(player.Hand, deck[0], deck[1])
		dealer.Hand = append(dealer.Hand, deck[2], deck[3])
		deck = deck[4:]
		
		fmt.Printf("\n%s's hand: ", player.Name)
		displayHand(player.Hand, false)
		fmt.Printf("Dealer's hand: ")
		displayHand(dealer.Hand, true)
		
		// Player's turn
		for {
			player.Score = calculateScore(player.Hand)
			if player.Score > 21 {
				break
			}
			
			fmt.Printf("\nYour score: %d\n", player.Score)
			fmt.Print("Do you want to (H)it or (S)tand? ")
			choice, _ := reader.ReadString('\n')
			choice = strings.ToUpper(strings.TrimSpace(choice))
			
			if choice == "H" {
				player.Hand = append(player.Hand, deck[0])
				deck = deck[1:]
				fmt.Printf("\n%s's hand: ", player.Name)
				displayHand(player.Hand, false)
			} else if choice == "S" {
				break
			}
		}
		
		// Dealer's turn
		fmt.Printf("\nDealer's hand: ")
		displayHand(dealer.Hand, false)
		
		dealer.Score = calculateScore(dealer.Hand)
		for dealer.Score < 17 {
			dealer.Hand = append(dealer.Hand, deck[0])
			deck = deck[1:]
			dealer.Score = calculateScore(dealer.Hand)
			fmt.Printf("Dealer hits: ")
			displayHand(dealer.Hand, false)
		}
		
		// Determine winner
		player.Score = calculateScore(player.Hand)
		dealer.Score = calculateScore(dealer.Hand)
		
		fmt.Printf("\nFinal scores:\n")
		fmt.Printf("%s: %d\n", player.Name, player.Score)
		fmt.Printf("Dealer: %d\n", dealer.Score)
		
		if player.Score > 21 {
			fmt.Println("Bust! Dealer wins!")
		} else if dealer.Score > 21 {
			fmt.Printf("%s wins!\n", player.Name)
		} else if player.Score > dealer.Score {
			fmt.Printf("%s wins!\n", player.Name)
		} else if dealer.Score > player.Score {
			fmt.Println("Dealer wins!")
		} else {
			fmt.Println("It's a tie!")
		}
		
		fmt.Print("\nPlay again? (Y/N): ")
		again, _ := reader.ReadString('\n')
		again = strings.ToUpper(strings.TrimSpace(again))
		if again != "Y" {
			break
		}
	}
	
	fmt.Println("Thanks for playing!")
}
