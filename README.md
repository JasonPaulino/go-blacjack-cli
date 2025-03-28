# CLI Blackjack

A simple command-line implementation of Blackjack in Go.

## Features
- Text-based user interface with card suit symbols
- Player vs Dealer gameplay
- Standard Blackjack rules
- Ace handling (1 or 11 points)
- Option to play multiple games
- Colorful interface with animations
- ASCII art logo
- Card dealing animations
- Colored suits (red hearts/diamonds, cyan spades/clubs)
- Score display for both player and dealer
- Continuous gameplay with option to quit anytime

## Installation

To install the game globally, run:

```bash
go install github.com/JasonPaulino/go-blacjack-cli/cmd/blackjack@latest
```

After installation, you can run the game from anywhere by typing:

```bash
blackjack
```

## Development

To run the game locally during development:

```bash
go run cmd/blackjack/main.go
```

## How to Play
1. Enter your name when prompted
2. You'll be dealt two cards
3. Choose to:
   - (H)it: Get another card
   - (S)tand: Keep your current hand
   - (Q)uit: Exit the game
4. Try to get as close to 21 as possible without going over
5. Beat the dealer's hand to win!
6. Press Enter to play again or 'Q' to quit after each round
