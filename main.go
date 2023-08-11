package main

import (
	"bufio"
	"fmt"
	"os"
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
}

func containsPoint(points []Point, p Point) bool {
	return true
}

func placeFood() {
}

func update() {

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
