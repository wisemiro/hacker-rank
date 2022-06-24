package algorithims

// Given a square matrix, calculate the absolute difference between the sums of its diagonals.

// For example, the square matrix

// is shown below:

// 1 2 3
// 4 5 6
// 9 8 9

// The left-to-right diagonal = 1 + 5 + 9 = 15
// . The right to left diagonal = 3 + 5 + 9 = 17
// Their absolute difference is 15 - 17 = 2

// Function description

// Complete the

// function in the editor below.

// diagonalDifference takes the following parameter:

//     int arr[n][m]: an array of integers

// Return

//     int: the absolute diagonal difference

// Input Format

// The first line contains a single integer,
// , the number of rows and columns in the square matrix .
// Each of the next lines describes a row, , and consists of space-separated integers

// .

// Constraints

// Output Format

// Return the absolute difference between the sums of the matrix's two diagonals as a single integer.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	var rDiagonal int32
	var lDiagonal int32
	lenArr := int32(len(arr))

	for row := int32(0); row < lenArr; row++ {
		rDiagonal += arr[row][row]
		lDiagonal += arr[row][int32(len(arr[row]))-1-row]
	}
	if lDiagonal > rDiagonal {
		return lDiagonal - rDiagonal
	} else {
		return rDiagonal - lDiagonal
	}
	// sol2
	// for i := range arr{
	//     lDiagonal += arr[i][i]
	//     rDiagonal += arr[i][int32(len(arr[i])) - 1 - int32(i)]
	//     i ++
	// }
	// ans := rDiagonal - lDiagonal
	// return ans
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
