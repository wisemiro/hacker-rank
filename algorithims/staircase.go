package algorithims

// Staircase detail

// This is a staircase of size

// :

//    #
//   ##
//  ###
// ####

// Its base and height are both equal to

// . It is drawn using # symbols and spaces. The last line is not preceded by any spaces.

// Write a program that prints a staircase of size

// .

// Function Description

// Complete the staircase function in the editor below.

// staircase has the following parameter(s):

//     int n: an integer

// Print

// Print a staircase as described above.

// Input Format

// A single integer,

// , denoting the size of the staircase.

// Constraints

// .

// Output Format

// Print a staircase of size

// using # symbols and spaces.

// Note: The last line must have

// spaces in it.

// Sample Input

// 6

// Sample Output

//      #
//     ##
//    ###
//   ####
//  #####
// ######

// Explanation

// The staircase is right-aligned, composed of # symbols and spaces, and has a height and width of n = 6

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func staircase(n int32) {
	// Write your code here
	for i := 1; int32(i) <= n; i++ {
		for j := 1; j <= int(n)-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("#") //notice use of print not println
		}
		fmt.Println()
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
