// Example of Cellular Automaton Rule 30
// The bit pattern of 30 is "00011110"
// So the 4th, 5th, 6, and 7th patterns get a "1" below their middle value.
// "111" "110" "101" "100" "011" "010" "001" "000"
//   0     0     0     1     1     1     1     0
// This continues until the bottom of the image
// graphical representation of rule 30
// http://mathworld.wolfram.com/images/eps-gif/ElementaryCARule030_1000.gif
//
// interesting cellular automatons
// [30 45 57 60 67 73 90 91 107 110 124 129 131 135 137 147 150]
//
// http://www.pheelicks.com/2013/10/intro-to-images-in-go-part-1/

//	bugfix 8-mar-2016
//	defer outFile.Close() changed to outFile.Close() 
// 	"too many open files" error
//	deferred calls are only executed when the function exits. 
// 	source: https://groups.google.com/forum/#!topic/golang-nuts/7yXXjgcOikM

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"fmt"
	"time"
)

type myColor color.RGBA

func main() {

	t1 := time.Now()

	interestingRules := []int{30,45,57,60,67,73,90,91,107,110,124,129,131,135,137,147,150}
	for _, rule := range interestingRules {

//	for rule := 0; rule <= 255; rule+=1 {	
		img := makeCellularAutomaton(rule)

		// Output file	
		outFilename := fmt.Sprintf("cellularAutomaton%v.png", rule)
		outFile, err := os.Create(outFilename)
		if err != nil {
			log.Fatal(err)
		}
		
		log.Print("Saving image to: ", outFilename)
		png.Encode(outFile, img)
		
		outFile.Close() // removed defer - "too many open files" error
	}
	
	fmt.Println( time.Since(t1) )

}

func makeCellularAutomaton( ruleNumber int ) image.Image {

	// Make yellow image
	yellow := color.RGBA{ 255, 255, 0, 255 }
	orange := color.RGBA{ 255, 150,10,255 }

	width, height := 1000,500
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			img.Set(x, y, yellow)
		}
	}
	img.Set(int(size.X / 2),0, orange)


	// Make 8 bit pattern	of number e.g. 2 => "00000010", 
	b8 := []byte(strconv.FormatUint(uint64(ruleNumber), 2))
	for len(b8)<8 {
		b8 = append([]byte{'0'}, b8...)
	}
	// fmt.Println(string(b8))


	// Make Rule
	rule := []color.RGBA{} 
	for i:= 0; i<8; i+=1 {
		if b8[i] == '1' {
			b3 := []byte(strconv.FormatUint(uint64(i), 2))
			for len(b3)<3 {
				b3 = append([]byte{'0'}, b3...)
			}
			// The bit patterns generated here (the b3s) are inverted
			// that is 001 instead of 110 and so on
			// It would be nice to fix this since it makes debugging unclear
			// fmt.Println(string(b3))

			for _, char := range b3 {
				if char == '0' {
					rule = append(rule, orange)
				} else {
					rule = append(rule, yellow)
				}
			}
		}
	}
	
//	for _, col := range rule {
//		fmt.Println(rule[0] == orange)
//		fmt.Print(myColor(col), " ")
//	}

	// Generate Cellular Automaton
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			for i := 0; i<len(rule)-2; i+=3 {
				if rule[i] == img.At(x,y) && rule[i+1] == img.At(x+1,y) && rule[i+2] == img.At(x+2,y) {
					img.Set(x+1, y+1, orange)
				}
			}
		}
	}
	
	return img
}

func ( c myColor ) String() string {
	return fmt.Sprintf("%v.%v.%v", c.R, c.G, c.B)
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}