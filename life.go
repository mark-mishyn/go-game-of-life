package main

import (
	"fmt"
	"math/rand"
	"time"
	"os/exec"
	"os"
	"strings"
)

func getRandBool() bool{
	return rand.Intn(2) == 1
}

func clearScreen(){
	clearScreenCommand := exec.Command("clear")
	clearScreenCommand.Stdout = os.Stdout
	clearScreenCommand.Run()
}

func generateInitial(size int) [][]bool{
	initialMatrix :=  [][]bool{}
	for x:=0; x<size; x++{
		row := []bool{}
		for y:=0; y<size; y++{
			row = append(row, getRandBool())
		}
		initialMatrix = append(initialMatrix, row)
	}

	return initialMatrix
}

func printMatrix(size int, matrix [][]bool){
	fmt.Println(strings.Repeat("-", size + 4))
	for _, row := range matrix{
		stringRow := ""
		for _, value:= range row{
			if value{
				stringRow += "*"
			} else {
				stringRow += " "
			}
		}
		fmt.Println("|", stringRow, "|")
	}
	fmt.Println(strings.Repeat("-", size + 4))
}

func getNeighborsCoordinates(x, y int) [][]int{
	return [][]int{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
		{x + 1, y + 1},
		{x - 1, y - 1},
		{x + 1, y - 1},
		{x - 1, y + 1},
	}
}

func normalizeNeighborsCoordinates(size, x, y int) [][]int{
	normalizedCoordinates:= [][]int{}

	for _, coord := range getNeighborsCoordinates(x,y){
		xCoord := coord[0]
		yCoord := coord[1]
		if xCoord == size{
			xCoord = 0
		} else if xCoord == -1{
			xCoord = size - 1
		}
		if yCoord == size{
			yCoord = 0
		} else if yCoord == -1{
			yCoord = size - 1
		}

		normalizedCoordinates = append(normalizedCoordinates, []int{xCoord, yCoord})
	}

	return normalizedCoordinates
}

func getAliveNeighborsCount(size, x, y int, matrix [][]bool) int{
	sum := 0
	for _, coord := range normalizeNeighborsCoordinates(size, x, y){
		xCoord := coord[0]
		yCoord := coord[1]
		if matrix[xCoord][yCoord]{
			sum++
		}
	}
	return sum
}

func getNewGeneration(size int, matrix [][]bool) [][]bool{
	new_matrix := [][]bool{}
	for i:=0; i<size; i++{
		row := []bool{}
		for j:=0; j<size; j++{
			isAlive := matrix[i][j]
			neighborsCount := getAliveNeighborsCount(size, i, j, matrix)

			var newState bool
			switch  {
				case isAlive && neighborsCount < 2:
					newState = false
				case isAlive && (neighborsCount == 2 || neighborsCount == 3):
					newState = true
				case isAlive && neighborsCount > 3:
					newState = false
				case isAlive == false && neighborsCount == 3:
					newState = true
				default:
					newState = isAlive
			}
			row = append(row, newState)
		}
		new_matrix = append(new_matrix, row)
	}
	return new_matrix
}

func main() {
	size := 45
	newMatrix := generateInitial(size)
	printMatrix(size, newMatrix)
	for {
		time.Sleep(time.Millisecond * 150)
		clearScreen()
		oldMatrix := newMatrix
		newMatrix = getNewGeneration(size, oldMatrix)
		printMatrix(size, newMatrix)
	}
}