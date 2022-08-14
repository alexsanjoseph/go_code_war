package kata

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func stringSum(x string) int {
	splitVal := strings.Split(x, "")
	sum := 0
	for _, x := range splitVal {
		xInt, err := strconv.ParseInt(x, 10, 32)
		if err != nil {
			panic("Why error?")
		}
		sum += int(xInt)
	}
	return sum
}

func OrderWeight(input string) string {
	fmt.Println(input)

	// Convert to slice
	splitString := strings.Split(input, " ")
	// splitInt := make([]int, len(splitString))

	// for i, x := range splitString {
	// 	xInt, err := strconv.ParseInt(x, 10, 32)
	// 	if err != nil {
	// 		panic("Why error?")
	// 	}
	// 	splitInt[i] = int(xInt)
	// }
	// fmt.Println(splitInt)
	sort.Slice(splitString, func(i, j int) bool {
		iSum := stringSum(splitString[i])
		jSum := stringSum(splitString[j])
		if iSum == jSum {
			return splitString[i] < splitString[j]
		}
		return iSum < jSum
	})
	fmt.Println(splitString)

	// sort using comparator

	return strings.Join(splitString, " ")
}
