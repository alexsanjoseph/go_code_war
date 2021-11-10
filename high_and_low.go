package kata

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	splitString := strings.Split(in, " ")
	smallest := math.MaxInt
	largest := math.MinInt
	for _, val := range splitString {
		conv, _ := strconv.Atoi(val)
		if conv < smallest {
			smallest = conv
		}
		if conv > largest {
			largest = conv
		}
	}
	return fmt.Sprintf("%d %d", largest, smallest)
}
