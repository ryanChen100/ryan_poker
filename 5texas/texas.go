package main

import (
	"fmt"
	"sort"
	"time"
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

type cardType int

// String
func (state cardType) String() string {
	return [...]string{
		"",
		"皇家同花顺", //1
		"同花顺",   //2
		"金刚",    //3
		"葫芦",    //4
		"同花",    //5
		"顺子",    //6
		"三条",    //7
		"两对",    //8
		"一对",    //9
		"高牌",    //10

		"END",
	}[state]
}

var allCard = []int{0x102, 0x103, 0x104, 0x105, 0x106, 0x107, 0x108, 0x109, 0x10a, 0x10b, 0x10c, 0x10d, 0x10e,
	0x202, 0x203, 0x204, 0x205, 0x206, 0x207, 0x208, 0x209, 0x20a, 0x20b, 0x20c, 0x20d, 0x20e,
	0x302, 0x303, 0x304, 0x305, 0x306, 0x307, 0x308, 0x309, 0x30a, 0x30b, 0x30c, 0x30d, 0x30e,
	0x402, 0x403, 0x404, 0x405, 0x406, 0x407, 0x408, 0x409, 0x40a, 0x40b, 0x40c, 0x40d, 0x40e}

// combinations 取所有組合
func combinations(arr []int, n int) [][]int {
	var helper func([]int, int, int)
	res := [][]int{}
	data := make([]int, n)

	helper = func(arr []int, n int, idx int) {
		if idx == n {
			temp := make([]int, n)
			copy(temp, data)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(arr); i++ {
			data[idx] = arr[i]
			helper(arr[i+1:], n, idx+1)
		}
	}
	helper(arr, n, 0)
	return res
}

func main() {
	// data := [][]string{
	// 	{"花色", "點數", "牌型"},
	// }
	// for _, combination := range combinations(allCard, 7) {
	// 	flush := ""
	// 	point := ""
	// 	for _, handCard := range combination {
	// 		flush += strconv.Itoa(getHighestDigit(handCard)) + " "
	// 		point += strconv.Itoa(getLowestDigit(handCard)) + " "
	// 	}

	// 	data = append(data, []string{flush, point, cardType(f(combination[:2], combination[2:])).String()})

	// }
	// fmt.Println("=================")
	// create_file.CreateCsv(data)
	start := time.Now()
	fmt.Println(f([]int{0x10b, 0x10a}, []int{0x105, 0x104, 0x10c, 0x10d, 0x10e}))
	fmt.Println(time.Since(start))
}

func f(self, board []int) int {
	newCard := append(self, board...)
	cardType := repeat(newCard)
	straightCard := straight(newCard)

	if cardType > straightCard {
		return straightCard
	}
	return cardType
}

func repeat(input []int) int {
	statistics := make(map[int]int)
	for _, v := range input {
		statistics[getLowestDigit(v)]++
	}

	appearAgain := ""
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
	checkDuplicates := removeDuplicates(sortInput)

	recodeStraight := []int{sortInput[0]}
	if len(checkDuplicates) >= 5 {

		for i := 1; i < len(sortInput); i++ {
			if sortInput[i] != sortInput[i-1]+1 {
				if sortInput[i] == 14 && sortInput[i-1] == 5 {
					sortInput[i] = 1
					recodeStraight = append(recodeStraight, sortInput[i])
					continue
				}

				if len(checkDuplicates)-i >= 5 {
					recodeStraight = []int{sortInput[i]}
					continue
				}

				if len(recodeStraight) >= 5 {
					break
				}

				isStraight = false
				break
			}
			recodeStraight = append(recodeStraight, sortInput[i])
		}
	} else {
		isStraight = false
	}

	isFlush := flush(input)

	if isStraight && isFlush {
		return calculatorStraight(input)
	} else if isFlush {
		return 5
	} else if isStraight {
		return 6
	} else {
		return cardType
	}

}

func flush(input []int) bool {
	flushMap := make(map[int]int)
	for _, v := range input {
		flushMap[getHighestDigit(v)]++
	}

	for _, v := range flushMap {
		if v >= 5 {
			return true
		}
	}
	return false
}

func calculatorStraight(input []int) int {
	sort.Ints(input)
	recodeStraight := []int{input[0]}
	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1]+1 {
			if getLowestDigit(input[i]) == 14 && getLowestDigit(input[i-1]) == 5 {
				input[i] = 1
				recodeStraight = append(recodeStraight, input[i])
				continue
			}

			if len(input)-i >= 5 {
				recodeStraight = []int{input[i]}
				continue
			}

			if len(recodeStraight) >= 5 {
				break
			}

			return 5
		}
		recodeStraight = append(recodeStraight, input[i])
	}

	sort.Ints(recodeStraight)
	if getLowestDigit(recodeStraight[len(recodeStraight)-1]) == 14 {
		return 1
	} else {
		return 2
	}
}

func removeDuplicates(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	result := slice[:1]
	for _, v := range slice {
		if v != result[len(result)-1] {
			result = append(result, v)
		}
	}

	return result
}

func getLowestDigit(hexNumber int) int {
	return hexNumber & 0xF
}

func getHighestDigit(hexNumber int) int {
	firstDigit := hexNumber >> 8
	firstDigit = firstDigit & 0xF

	return firstDigit
}
