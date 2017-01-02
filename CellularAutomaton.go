// Graphical representation of Rule 30
// http://mathworld.wolfram.com/images/eps-gif/ElementaryCARule030_1000.gif
//
//  Example of Cellular Automaton Rule 30
// The bit pattern of 30 is "00011110"
// So the 4th, 5th, 6, and 7th patterns get a "1" below their middle value.
// "111" "110" "101" "100" "011" "010" "001" "000"
//   0     0     0     1     1     1     1     0
// This continues until the bottom of the image
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
//
// 1-1-2017
// Tried using goroutines to speed it up with some success.
// At first I just put the writePNG function in a goroutine.
// This reduced execution time to about 75% of what it was which was nice.
// After this I tried putting the whole thing into goroutines using the sync library
// This proved even better even on a macbook pro with only 2 cores.
//
// Example of sync waitgroup usage example of goroutines that I used
// http://stackoverflow.com/questions/19208725/example-for-sync-waitgroup-correct

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	t1 := time.Now()

	if true {
		// black := color.RGBA{0, 0, 0, 255}
		// red := color.RGBA{255, 0, 0, 255}
		yellow := color.RGBA{255, 255, 0, 255}
		orange := color.RGBA{255, 150, 10, 255}
		width := 1000 // height is half width

		var wg sync.WaitGroup
		for _, rule := range []int{30, 45, 57, 60, 67, 73, 90, 91, 107, 110, 124, 129, 131, 135, 137, 147, 150} {
			// for rule := 0; rule <= 255; rule += 1 {
			wg.Add(1)
			go goRoutine(rule, width, yellow, orange, &wg)
			// img := makeCellularAutomaton(rule, width, yellow, orange)      // height is half width
			// go writePNG(fmt.Sprintf("cellularAutomaton%v.png", rule), img) // goroutine makes it significantly faster
		}
		wg.Wait()
	}

	fmt.Println(time.Since(t1))
	fmt.Println("Done")
}

func goRoutine(rule, width int, yellow, orange color.RGBA, wg *sync.WaitGroup) {

	img := makeCellularAutomaton(rule, width, yellow, orange)   // height is half width
	writePNG(fmt.Sprintf("cellularAutomaton%v.png", rule), img) // goroutine makes it significantly faster
	wg.Done()

}

func writePNG(outFilename string, Img image.Image) {

	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, Img)

	outFile.Close() // removed defer - "too many open files" error
}

func makeCellularAutomaton(RuleNumber, width int, color1, color2 color.RGBA) image.Image {

	Img := image.NewRGBA(image.Rect(0, 0, width, int(width/2)))
	ImgSize := Img.Bounds().Size()
	Rule := []color.RGBA{} // 3 pixel pattern to match

	// Color in image - fill with color1 and put a dot top center of color2
	if true {
		for x := 0; x < ImgSize.X; x++ {
			for y := 0; y < ImgSize.Y; y++ {
				Img.Set(x, y, color1)
			}
		}
		Img.Set(int(ImgSize.X/2), 0, color2)
	}

	// Make Rule - Make a set of 3 pixel patterns that will be used to generate the image
	if true {

		// Make 8 bit pattern of RuleNumber e.g. 2 => "00000010",
		b8 := []byte(strconv.FormatUint(uint64(RuleNumber), 2))
		for len(b8) < 8 {
			b8 = append([]byte{'0'}, b8...)
		}

		// Cellular Automaton Rule 30 Example
		// The bit pattern of 30 is "00011110"
		// so the 4th, 5th, 6, and 7th patterns get a "1" below their middle value.
		// "111" "110" "101" "100" "011" "010" "001" "000"
		//   0     0     0     1     1     1     1     0
		// These are what will define the rule
		// rule30 = ["100","011","010","001"]
		// where the zeroes and ones are actually different colored pixels.

		for i := 0; i < 8; i += 1 {
			if b8[7-i] == '1' { // since numbers grow from right to left
				b3 := []byte(strconv.FormatUint(uint64(i), 2))
				for len(b3) < 3 {
					b3 = append([]byte{'0'}, b3...)
				}
				for _, char := range b3 {
					if char == '0' {
						Rule = append(Rule, color1)
					} else {
						Rule = append(Rule, color2)
					}
				}
			}
		}
	}

	// Draw Cellular Automaton Image
	for y := 0; y < ImgSize.Y; y++ {
		for x := 0; x < ImgSize.X; x++ {
			for i := 0; i < len(Rule)-2; i += 3 {
				if Rule[i] == Img.At(x, y) && Rule[i+1] == Img.At(x+1, y) && Rule[i+2] == Img.At(x+2, y) {
					Img.Set(x+1, y+1, color2)
				}
			}
		}
	}

	return Img
}

type myColor color.RGBA

func (c myColor) String() string {
	return fmt.Sprintf("%v.%v.%v", c.R, c.G, c.B)
}
