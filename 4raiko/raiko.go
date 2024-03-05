package main

import (
	"fmt"
)

// 1、皇家同花顺：如果花色一样，数字分别是10,J,Q,K,A
// 2、同花顺：如果花色一样，数字是连续的，皇家同花顺除外
// 3、金刚：其中4张牌数字一样
// 4、葫芦：其中3张牌数字一样，另外2张牌数字一样
// 5、同花：花色一样，数字不连续
// 6、顺子：数字是连续，花色不一样
// 7、三条：其中3张牌数字一样，另外2张牌数字不一样
// 8、两对：其中2张牌数字一样，另外其中2张牌数字一样，最后一张数字不一样
// 9、一对：其中2张牌数字一样，另外数字不一样
// 10、高牌：什么都不是
// 0x50f代表小王
// 0x610代表大王

var test1 = []int{0x102, 0x103, 0x104, 0x105, 0x106, 0x107, 0x108, 0x109, 0x10a, 0x10b, 0x10c, 0x10d, 0x10e}
var test2 = []int{0x202, 0x203, 0x204, 0x205, 0x206, 0x207, 0x208, 0x209, 0x20a, 0x20b, 0x20c, 0x20d, 0x20e}
var test3 = []int{0x302, 0x303, 0x304, 0x305, 0x306, 0x307, 0x308, 0x309, 0x30a, 0x30b, 0x30c, 0x30d, 0x30e}
var test4 = []int{0x402, 0x403, 0x404, 0x405, 0x406, 0x407, 0x408, 0x409, 0x40a, 0x40b, 0x40c, 0x40d, 0x40e}

func main() {

	fmt.Println(f([]int{0x10a, 0x10b, 0x10c, 0x10d, 0x10e}))

}

// var raiko = []int{0x50f, 0x610}

func f(input []int) int {
	if len(input) != 5 {
		input = input[:5]
	}

	cardType := 10
	statistics := make(map[int]int)
	raikoCount := 0
	samePoint := []int{}

	for _, v := range input {
		if v == 0x50f || v == 0x610 {
			raikoCount++
		} else {
			statistics[getLowestDigit(v)]++
		}
	}

	for _, v := range statistics {
		if v > 1 {
			samePoint = append(samePoint, v)
		}
	}

	cardType = calculatorSame(raikoCount, samePoint)

	cardType = calculatorStraight(raikoCount, input, statistics)

	return cardType
}

func calculatorSame(raiko int, same []int) int {
	if len(same) == 0 && raiko == 0 {
		return 10
	}

	max := raiko
	for _, v := range same {
		v += raiko
		if v > max {
			max = v
		}
	}

	if max == 4 {
		return 3
	} else if max == 3 && len(same) > 1 {
		return 4
	} else if max == 3 {
		return 7
	} else if max == 2 && len(same) > 1 {
		return 8
	} else if max == 2 {
		return 9
	}

	return 10
}

func calculatorStraight(raiko int, input []int, statistics map[int]int) int {
	cardType := 10

	switch raiko {
	case 0:
		cardType = straight(input)
	case 1:
		cardType = oneRaikoStraight(input)
	case 2:
		cardType = twoRaikoStraight(input)
	}

	return cardType

}

func oneRaikoStraight(input []int) int {
	isStraight := true
	cardType := 10
	for i := 0; i < len(input)/2-1; i++ {
		if input[i]+input[len(input)-1-i] != input[len(input)/2]*2 {
			isStraight = false
			break
		}
	}

	isFlush := flush(input)

	if isStraight && isFlush {
		if input[len(input)/2] == 0x10c {
			return 1
		} else {
			return 2
		}
	} else if isStraight {
		return 6
	} else if isFlush {
		return 5
	} else {
		return cardType
	}

}

func twoRaikoStraight(input []int) int {
	isStraight := true
	cardType := 10
	for i := 0; i < len(input)/2-1; i++ {
		if input[i]+input[len(input)-1-i] != input[len(input)/2]*2 {
			isStraight = false
			break
		}
	}

	isFlush := flush(input)

	if isStraight && isFlush {
		if input[len(input)/2] == 0x10c {
			return 1
		} else {
			return 2
		}
	} else if isStraight {
		return 6
	} else if isFlush {
		return 5
	} else {
		return cardType
	}

}

func straight(input []int) int {
	isStraight := true
	cardType := 10

	// [0x10e,0x102,0x103,0x104,0x105]例外排除
	// 包含 0x10e,0x102
	for i := 0; i < len(input)/2-1; i++ {
		if getHighestDigit(input[i])+getHighestDigit(input[len(input)-1-i]) != getHighestDigit(input[len(input)/2])*2 {
			isStraight = false
			break
		}
	}

	isFlush := flush(input)

	if isStraight && isFlush {
		if input[len(input)/2] == 0x10c {
			return 1
		} else {
			return 2
		}
	} else if isStraight {
		return 6
	} else if isFlush {
		return 5
	} else {
		return cardType
	}

}

func flush(input []int) bool {
	isFlush := true
	for i := 1; i < len(input); i++ {
		if getHighestDigit(input[i]) != getHighestDigit(input[i-1]) {
			isFlush = false
			return isFlush
		}
	}
	return isFlush
}

func getLowestDigit(hexNumber int) int {
	// 使用按位与操作获取最低位的值
	// lowestDigit := hexNumber & 0xF

	return hexNumber & 0xF
}

func getHighestDigit(hexNumber int) int {
	// 右移 12 位，将最高位移到最右边
	firstDigit := hexNumber >> 8

	// 使用按位与操作获取最高位的值
	firstDigit = firstDigit & 0xF

	return firstDigit
}
