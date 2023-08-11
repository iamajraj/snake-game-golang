package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	width  = 20
	height = 10
)

type Point struct {
	x, y int
}

type Snake struct {
	head      Point
	body      []Point
	direction Point
}

var (
	food     Point
	snake    Snake
	gameOver bool
)

func setup() {
	snake = Snake{
		head:      Point{width / 2, height / 2},
		direction: Point{1, 0},
	}
	placeFood()
	gameOver = false
}

func draw() {
	clearScreen()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == snake.head.x && y == snake.head.y {
				fmt.Print("■") // Snake head
			} else if containsPoint(snake.body, Point{x, y}) {
				fmt.Print("o") // Snake body
			} else if x == food.x && y == food.y {
				fmt.Print("🍏") // Food emoji
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen
	}
}

func containsPoint(points []Point, p Point) bool {
	for _, point := range points {
		if point == p {
			return true
		}
	}
	return false
}

func placeFood() {
	rand.Seed(time.Now().UnixNano())
	food = Point{rand.Intn(width), rand.Intn(height)}
}

func update() {
	newHead := Point{
		x: snake.head.x + snake.direction.x,
		y: snake.head.y + snake.direction.y,
	}

	snake.body = append([]Point{snake.head}, snake.body...)

	if newHead == food {
		placeFood()
	} else {
		snake.body = snake.body[:len(snake.body)-1]
	}

	snake.head = newHead

	if snake.head.x < 0 || snake.head.x >= width || snake.head.y < 0 || snake.head.y >= height {
		gameOver = true
		return
	}

	if containsPoint(snake.body, snake.head) {
		gameOver = true
		return
	}
}

func handleInput() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch char {
	case 'w':
		snake.direction = Point{0, -1}
	case 'a':
		snake.direction = Point{-1, 0}
	case 's':
		snake.direction = Point{0, 1}
	case 'd':
		snake.direction = Point{1, 0}
	}
}

func main() {
	setup()

	for !gameOver {
		draw()
		handleInput()
		update()
	}

	fmt.Println("Game Over!")
}
