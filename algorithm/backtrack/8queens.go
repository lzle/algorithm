package backtrack

import (
	"fmt"
	"os"
)

var result [8]int

func call8Queens(row int) {
	if row == 8 {
		printQueues()
		os.Exit(100)
	}
	for column:=0; column<8; column++ {
		if isOk(row,column) == true {
			result[row] = column
			call8Queens(row+1)
		}
	}
}

func isOk(row int, column int) bool {
	left,right := column-1,column+1
	for i:=row-1; i>=0; i-- {
		if result[i] == column {
			return false
		}
		if left >= 0 {
			if result[i] == left {
				return false
			}
		}
		if right < 8 {
			if result[i] == right{
				return false
			}
		}
		left--
		right++
	}
	return true
}

func printQueues()  {
	for row:=0; row<8; row++ {
		for column:=0; column<8; column++ {
			if result[row] == column {
				fmt.Printf("Q ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
}


/*
func main()  {
	for i:=0; i<8; i++ {
		result[i] = -1
	}
	call8Queens(0)
}
 */