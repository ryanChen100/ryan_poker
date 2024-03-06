package main

import (
	"fmt"
	"sort"
)

// 请完成函数f，输入的5个数字代表5张牌，含义如下：
// 0x102,0x103,0x104,0x105,0x106,0x107,0x108,0x109,0x10a,0x10b,0x10c,0x10d,0x10e分别代表方块2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x202,0x203,0x204,0x205,0x206,0x207,0x208,0x209,0x20a,0x20b,0x20c,0x20d,0x20e分别代表梅花2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x302,0x303,0x304,0x305,0x306,0x307,0x308,0x309,0x30a,0x30b,0x30c,0x30d,0x30e分别代表红桃2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x402,0x403,0x404,0x405,0x406,0x407,0x408,0x409,0x40a,0x40b,0x40c,0x40d,0x40e分别代表黑桃2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x50f代表小王
// 0x610代表大王
// 小王大王可以变为任意牌，要求算出小王大王变换后最大牌型
// 返回的数字含义如下：
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
	0x402, 0x403, 0x404, 0x405, 0x406, 0x407, 0x408, 0x409, 0x40a, 0x40b, 0x40c, 0x40d, 0x40e, 0x50f, 0x610}

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
	// for _, combination := range combinations(allCard, 5) {
	// 	flush := ""
	// 	point := ""
	// 	for _, handCard := range combination {
	// 		flush += strconv.Itoa(getHighestDigit(handCard)) + " "
	// 		point += strconv.Itoa(getLowestDigit(handCard)) + " "
	// 	}
	// 	tmp := f(combination)
	// 	data = append(data, []string{flush, point, cardType(tmp).String(), strconv.Itoa(tmp)})

	// }
	// create_file.CreateCsv(data)
	fmt.Println(f([]int{0x102, 0x103, 0x104, 0x105, 0x10e}))

}

func f(input []int) int {
	if len(input) != 5 {
		return 0
	}

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
		samePoint = append(samePoint, v)
	}

	cardType := raikoRepeat(raikoCount, samePoint)
	straightCard := raikoStraight(raikoCount, input)
	if cardType > straightCard {
		return straightCard
	}
	return cardType

}

func raikoRepeat(raiko int, same []int) int {
	if len(same) == 0 && raiko == 0 {
		return 10
	}

	max := []int{}
	for _, v := range same {
		max = append(max, v)
	}

	sort.Ints(max)

	max[len(max)-1] += raiko

	if max[len(max)-1] > 3 {
		return 3
	} else if max[len(max)-1] == 3 && max[len(max)-2] == 2 {
		return 4
	} else if max[len(max)-1] == 3 {
		return 7
	} else if max[len(max)-1] == 2 {
		return 9
	}
	return 10
}

// "金刚",      //3
// "葫芦",      //4
// "同花",      //5
// "顺子",      //6
// "三条",      //7
// "两对",      //8
// "一对",      //9
// "高牌",      //10

func raikoStraight(raikoCount int, input []int) int {
	isStraight := true
	cardType := 10
	sortInput := []int{}

	for _, v := range input {
		sortInput = append(sortInput, getLowestDigit(v))
	}

	sort.Ints(sortInput)
	checkDuplicates := removeDuplicates(sortInput)

	if len(checkDuplicates)+raikoCount == 5 {

		passCount := raikoCount
		for i := 1; i < len(sortInput); i++ {
			fmt.Println(sortInput[i], sortInput[i-1])
			if sortInput[i] != sortInput[i-1]+1 {
				if i == len(sortInput)-1 && sortInput[i] == 14 && sortInput[i-1] == 5 {
					sortInput[len(sortInput)-1] = 1
					continue
				}

				if passCount > 0 {
					if sortInput[i] == sortInput[i-1]+2 || i == len(sortInput)-1 && sortInput[i] == 14 && sortInput[i-1] == 4 {
						passCount--
						continue
					}
				}

				if passCount > 1 {
					if sortInput[i] == sortInput[i-1]+3 || i == len(sortInput)-1 && sortInput[i] == 14 && sortInput[i-1] == 3 {
						passCount -= 2
						continue
					}
				}

				isStraight = false
				break
			}
		}
	} else {
		isStraight = false
	}
	isFlush := flush(raikoCount, input)

	if isStraight && isFlush {
		if sortInput[len(sortInput)-1] == 14 {
			return 1
		} else {
			return 2
		}
	} else if isFlush {
		return 5
	} else if isStraight {
		return 6
	} else {
		return cardType
	}

}

func flush(raiko int, input []int) bool {
	flushMap := make(map[int]int)
	for _, v := range input {
		flushMap[getHighestDigit(v)]++
	}
	max := 0
	for _, v := range flushMap {
		if v > max {
			max = v
		}
	}

	if max+raiko >= 5 {
		return true
	} else {
		return false
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
