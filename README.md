# GoCellularAutomaton

<pre>
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
</pre>
