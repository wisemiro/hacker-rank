package algorithims

// In this challenge, you are required to calculate and print the sum of the elements in an array, keeping in mind that some of those integers may be quite large.

// Function Description

// Complete the aVeryBigSum function in the editor below. It must return the sum of all array elements.

// aVeryBigSum has the following parameter(s):

//     int ar[n]: an array of integers .

// Return

//     long: the sum of all array elements

// Input Format

// The first line of the input consists of an integer
// .
// The next line contains

// space-separated integers contained in the array.

// Output Format

// Return the integer sum of the elements in the array.

// Constraints

// Sample Input

// 5
// 1000000001 1000000002 1000000003 1000000004 1000000005

// Output

// 5000000015

// When we add several integer values, the resulting sum might exceed the above range. You might need to use long int C/C++/Java to store such sums. 

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
   var totalSum int64
   for _ ,x := range ar {
       totalSum += x
   }
   return totalSum
}

func main() {
    arrValues := []int64{1000000001,1000000002, 1000000003}
    aVeryBigSum(arrValues)
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ar []int64

    for i := 0; i < int(arCount); i++ {
        arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        ar = append(ar, arItem)
    }

    result := aVeryBigSum(ar)

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
