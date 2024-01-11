package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

func ReadPBM(filename string) (*PBM, error) {
	file, err := os.Open("image.pbm")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var magicNumber string
	scanner.Scan()
	magicNumber = scanner.Text()

	scanner.Scan()
	width, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	scanner.Scan()
	height, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	data := make([][]bool, height)
	for i := 0; i < height; i++ {
		data[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			scanner.Scan()
			pixelValue := scanner.Text()
			data[i][j] = pixelValue == "1"
		}
	}

	return &PBM{data, width, height, magicNumber}, nil
}

// Size returns the width and height of the image.
func (pbm *PBM) Size() (int, int) {
	return pbm.width, pbm.height
}

// At returns the value of the pixel at (x, y).
func (pbm *PBM) At(x, y int) bool {
	return pbm.data[y][x]
}

// Set sets the value of the pixel at (x, y).
func (pbm *PBM) Set(x, y int, value bool) {
	pbm.data[y][x] = value
}

// Save saves the PBM image to a file and returns an error if there was a problem.
func (pbm *PBM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintf(writer, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)

	for _, row := range pbm.data {
		for _, val := range row {
			if val {
				fmt.Fprintf(writer, "1 ")
			} else {
				fmt.Fprintf(writer, "0 ")
			}
		}
		fmt.Fprintln(writer)
	}

	return writer.Flush()
}

// Invert inverts the colors of the PBM image.
func (pbm *PBM) Invert() {
	for i := 0; i < pbm.height; i++ {
		for j := 0; j < pbm.width; j++ {
			pbm.data[i][j] = !pbm.data[i][j]
		}
	}
}

// Flip flips the PBM image horizontally.
func (pbm *PBM) Flip() {
	for i := 0; i < pbm.height; i++ {
		for j := 0; j < pbm.width/2; j++ {
			pbm.data[i][j], pbm.data[i][pbm.width-j-1] = pbm.data[i][pbm.width-j-1], pbm.data[i][j]
		}
	}
}

// Flop flops the PBM image vertically.
func (pbm *PBM) Flop() {
	for i := 0; i < pbm.height/2; i++ {
		pbm.data[i], pbm.data[pbm.height-i-1] = pbm.data[pbm.height-i-1], pbm.data[i]
	}
}

// SetMagicNumber sets the magic number of the PBM image.
func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
}

func main() {
	// Example usage
	image, err := ReadPBM("image.pbm")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	width, height := image.Size()
	fmt.Println("Image Size:", width, "x", height)

	fmt.Println("Value at (1, 1):", image.At(1, 1))

	image.Invert()
	image.Save("inverted_image.pbm")
}
