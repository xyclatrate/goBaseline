package main

import (
	"fmt"
)

func main() {

	var numberOfMonsters = 20

	var text = string("Go Language is suitable for high speed rapid developement and it's syntax seems engaging. ")
	var thanksMessage = string(" Thanks Ken Thompson and Dennis Ritchie . The world can never develop without the community ")
	var nIntegerNumber = int(87)
	var n8IntegerNumber = int8(8)
	var n16integerNumber = int16(8983)
	var n32IntegerNumber = int32(89748)
	var n64IntegerNumber = int64(857465)
	var un8IntegerNumber = uint(838273892)
	var reallyBigNumber = float64(847483748374837847384783748738947384789.836473864736473467)
	fmt.Println(text, thanksMessage, numberOfMonsters, nIntegerNumber, n8IntegerNumber, n16integerNumber, n32IntegerNumber, n64IntegerNumber, un8IntegerNumber, reallyBigNumber)

}
