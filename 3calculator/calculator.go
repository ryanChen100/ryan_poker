package main

import (
	"fmt"
	"sort"
)

// 1、皇家同花顺：如果花色一样，数字分别是10,J,Q,K,A
// 2、同花顺：如果花色一样，数字是连续的，皇家同花顺除外，例如[0x109,0x10a,0x10b,0x10c,0x10d],[0x10e,0x102,0x103,0x104,0x105]
// 3、ok金刚：其中4张牌数字一样
// 4、ok葫芦：其中3张牌数字一样，另外2张牌数字一样
// 5、同花：花色一样，数字不连续
// 6、顺子：数字是连续，花色不一样
// 7、ok三条：其中3张牌数字一样，另外2张牌数字不一样
// 8、ok两对：其中2张牌数字一样，另外其中2张牌数字一样，最后一张数字不一样
// 9、ok一对：其中2张牌数字一样，另外数字不一样
// 10、高牌：什么都不是

var test1 = []int{0x102, 0x103, 0x104, 0x105, 0x106, 0x107, 0x108, 0x109, 0x10a, 0x10b, 0x10c, 0x10d, 0x10e}
var test2 = []int{0x202, 0x203, 0x204, 0x205, 0x206, 0x207, 0x208, 0x209, 0x20a, 0x20b, 0x20c, 0x20d, 0x20e}
var test3 = []int{0x302, 0x303, 0x304, 0x305, 0x306, 0x307, 0x308, 0x309, 0x30a, 0x30b, 0x30c, 0x30d, 0x30e}
var test4 = []int{0x402, 0x403, 0x404, 0x405, 0x406, 0x407, 0x408, 0x409, 0x40a, 0x40b, 0x40c, 0x40d, 0x40e}

func main() {

	fmt.Println(f([]int{0x10e, 0x502d, 0x402, 0x303, 0x204}))

}

func f(input []int) int {
	if len(input) != 5 {
		input = input[:5]
	}

	cardType := 10
	cardType = straight(input)
	if cardType != 10 {
		return cardType
	}

	cardType = repeat(input)
	if cardType != 10 {
		return cardType
	}

	return cardType
}

func repeat(input []int) int {
	statistics := make(map[int]int)
	for _, v := range input {
		statistics[getLowestDigit(v)]++
	}

	appearAgain := ""
	// appearValue := 0
	for _, count := range statistics {
		if count > 0 {
			if count == 4 {
				return 3
			} else if count == 3 {
				appearAgain += "3"
			} else if count == 2 {
				appearAgain += "2"
			}
		}
	}

	switch appearAgain {
	case "32", "23":
		return 4
	case "3":
		return 7

	case "22":
		return 8

	case "2":
		return 9
	default:
		return 10
	}
}

func straight(input []int) int {
	isStraight := true
	cardType := 10
	sortInput := []int{}
	for _, v := range input {
		sortInput = append(sortInput, getLowestDigit(v))
	}
	
	sort.Ints(sortInput)
	fmt.Println(sortInput)
	// [0x10e,0x102,0x103,0x104,0x105]例外排除
	// 包含 0x10e,0x102

	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1]+1 {
			// fmt.Printf("%X,   %X\n", getLowestDigit(input[i]), getLowestDigit(input[i-1]))
			// fmt.Println(getLowestDigit(input[i]), getLowestDigit(input[i-1]))
			// fmt.Println(0x002, 0x00e)
			if input[i] == 14 && input[i-1] == 2 {
				continue
			} else {
				isStraight = false
				break
			}
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
