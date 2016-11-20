package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func main() {
	for _, arg := range os.Args[1:] {
		fnum, err := strconv.ParseFloat(arg, 64)
		num, err := strconv.ParseInt(arg, 0, 64)
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		if num == 2 {
			fmt.Println("Prime")
			return
		} else if num < 2{
			fmt.Println("Not Prime")
			return
		} else if num % 2 == 0 {
			fmt.Println("Not Prime")
			return
		}
		//var float_num float64 = num
		//var pointFive float64 = 0.5
		sqr := math.Sqrt(fnum)
		isqr :=  int(math.Ceil(sqr))
		for i := 3; i<= isqr; i+=2 {
			var i64 int64
			i64 = int64(i)
			if num%i64 == 0 {
				fmt.Println(i)
				fmt.Println("Not Prime")
				return
			}
		}
		fmt.Println("Prime")

	}
}
