package main

import (
	"fmt"
	"sort"
	"time"
)

// 同花大顺（皇家同花顺或最大同花顺，Royal Straight Flush）
// 同花色的A，K，Q，J和10。
// 平手牌：公牌开出同花大顺，则所有未盖牌的牌手平手平分筹码。
// 平分,且平點數,比同花色手牌

// 同花顺（Straight Flush）
// 五张同花色的连续数字牌。同时有同花顺时，数字最大者为赢家。
// 平手牌：公牌开出同花顺为最大时，则所有未盖牌的牌手平手平分筹码。
// 平分,且平點數,比同花色手牌

// 铁支（四条，Four of a kind）
// 其中四张是相同数字的扑克牌，第五张是剩下牌组中最大的一张牌。若有一家以上持有四条（公牌开出四条），则比较第五张牌（起脚牌），最大者为赢家。
// 平手牌：公牌开出四条时，最后一张杂牌（或称为kicker、次大牌、踢脚牌，一副牌型组合中剩下来没有用作凑牌型的牌，用于牌型相同时比大小）数字也相同时。
// 平分,且平點數,比手牌

// 葫芦（夫佬或满堂红，Full house）
// 由三张相同数字及任何两张其他相同数字的扑克牌组成，如果同时有多人拿到葫芦，三张相同数字中数字较大者为赢家。如果使用多副牌且三张牌都一样，则再比两张牌中数字较大者赢家。
// 平手牌：五张牌数字都一样，则平分彩池。

// 同花（Flush）
// 此牌由五张不按顺序但相同花色的扑克牌组成，如果不只一人有此牌组，则牌面数字最大的人赢得该局，如果最大数字相同，则由第二、第三、第四或者第五张牌来决定胜负。
// 平手牌：公牌的同花就是最大的同花牌型时，平分彩池。
// 比同花色大小

// 顺子（Straight）
// 此牌由五张连续数字扑克牌组成，如果不只一人有此牌组，则五张牌中数字最大的赢得此局，10-J-Q-K-A为最大的顺子，A-2-3-4-5为最小的顺子。
// 平手牌：如果五张牌数字都相同，平分彩池。

// 三条（Three of a kind）
// 由三张相同数字和两张不同数字的扑克牌组成，如果不只一人有此牌组，则三张牌中数字者最大赢得该局。如果使用多副牌且三张牌数字大小相同，则比较不同点数的两张牌中数字较大者，若相同时再比第五张，数字大的人赢。
// 平手牌：如果五张牌数字都相同，则平分彩池。
//比後兩張

// 两对（Two pair）
// 两对数字相同但两两不同的扑克和一张杂牌组成，共五张牌。
// 平手牌：如果不只一人持有此牌型，持有数字比较大的对子者为赢家，若较大数字对子相同，则比较小对子的数字，如果两对对子数字都相同，那么第五张牌（kicker）数字较大者赢。如果连第五张牌数字也相同，则平分彩池。
//比最後一張

// 对子（Pair）
// 由两张相同数字的扑克牌和另三张无法组成牌型的杂牌组成。
// 平手牌：如果不只一人持有此牌型，则持有较大数字对子者为赢家，如果对牌数字相同，则依序比较剩下的三张牌，数字最大者为赢家，如果五张牌都一样，则平分彩池。
// 比最後三張

// 乌龙（高牌或散牌, High card，No-pair，Zilch）
// 无法组成以上任一牌型的杂牌。
// 平手牌：如果不只一人抓到此牌，则比较数字最大者，如果数字最大的相同，则依序比较第二、第三、第四和第五大的，如果五张牌都相同，则平分彩池。
// 比所有牌

// 请完成函数f，输入是2个int的slice，第1个slice的长度是2，代表自己的手牌，第2个slice的长度是5，代表公共牌，数字含义如下：
// 0x102,0x103,0x104,0x105,0x106,0x107,0x108,0x109,0x10a,0x10b,0x10c,0x10d,0x10e分别代表方块2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x202,0x203,0x204,0x205,0x206,0x207,0x208,0x209,0x20a,0x20b,0x20c,0x20d,0x20e分别代表梅花2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x302,0x303,0x304,0x305,0x306,0x307,0x308,0x309,0x30a,0x30b,0x30c,0x30d,0x30e分别代表红桃2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x402,0x403,0x404,0x405,0x406,0x407,0x408,0x409,0x40a,0x40b,0x40c,0x40d,0x40e分别代表黑桃2,3,4,5,6,7,8,9,10,J,Q,K,A

// var allCard = []int{0x102, 0x103, 0x104, 0x105, 0x106, 0x107, 0x108, 0x109, 0x10a, 0x10b, 0x10c, 0x10d, 0x10e,
//
//	0x202, 0x203, 0x204, 0x205, 0x206, 0x207, 0x208, 0x209, 0x20a, 0x20b, 0x20c, 0x20d, 0x20e,
//	0x302, 0x303, 0x304, 0x305, 0x306, 0x307, 0x308, 0x309, 0x30a, 0x30b, 0x30c, 0x30d, 0x30e,
//	0x402, 0x403, 0x404, 0x405, 0x406, 0x407, 0x408, 0x409, 0x40a, 0x40b, 0x40c, 0x40d, 0x40e}
var allCard = []int{0x102, 0x103, 0x104, 0x105, 0x106, 0x107, 0x108, 0x109, 0x10a, 0x10b, 0x10c, 0x10d, 0x10e,
	0x202, 0x203, 0x204, 0x205, 0x206, 0x207, 0x208, 0x209, 0x20a, 0x20b, 0x20c, 0x20d, 0x20e,
	0x302, 0x303, 0x304, 0x305, 0x306, 0x307, 0x308, 0x309, 0x30a, 0x30b, 0x30c, 0x30d, 0x30e,
	0x402, 0x403, 0x404, 0x405, 0x406, 0x407, 0x408, 0x409, 0x40a, 0x40b, 0x40c, 0x40d, 0x40e}

// 从7张牌中找出组成最大牌型的5张牌
// 只有1个对手，算出自己的胜率百分比，只保留整数部分
// 请提供较高性能的方案

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

func main() {
	fmt.Println(f([]int{0x102, 0x103}, []int{0x104, 0x105, 0x106, 0x107, 0x108}))
}

type card struct {
	self     []int
	board    []int
	input    []int
	cardType int

	cardPoint    int
	cardSecPoint int

	isFlish bool
}

var (
	winPro     chan int
	totalRound int
)

func f(self, board []int) int {
	newCard := append(self, board...)
	userCard := card{
		self:  self,
		board: board,
		input: newCard,
	}
	userCard.getUserType()
	fmt.Println(userCard)
	fmt.Println("======================")
	allCard = remove(allCard, newCard)
	fmt.Println(allCard)

	allPro := combinations(allCard, 2)
	totalRound = len(allPro)
	fmt.Println(totalRound)
	time.Sleep(5 * time.Second)
	saveLog := make(chan card, totalRound)
	winPro = make(chan int)
	go getLog(userCard, saveLog)
	fmt.Println(allPro)
	for _, v := range allPro {
		newPro := append(v, board...)
		pro := card{
			self:  v,
			board: board,
			input: newPro,
		}

		go func(pro card, saveLog chan card) {
			setLog(pro, saveLog)
		}(pro, saveLog)

	}
	sum := <-winPro
	return sum
}

func getLog(userCard card, saveLog chan card) {
	count := 0
	for log := range saveLog {
		if userCard.cardType > log.cardType {
			count++
		} else if userCard.cardType == log.cardType {
			if userCard.cardPoint > log.cardPoint {
				count++
			} else if userCard.cardPoint == log.cardPoint && userCard.cardSecPoint > log.cardSecPoint {
				count++
			} else if userCard.cardPoint == log.cardPoint && userCard.cardSecPoint == log.cardSecPoint {
				if finalCalculate(userCard, log) {
					count++
				}
			}
		}

		if count == totalRound {
			break
		}
	}

	winPro <- (count / totalRound) * 100

}

func finalCalculate(userCard, log card) bool {
	// 同花大顺（皇家同花顺或最大同花顺，Royal Straight Flush）
	// 同花色的A，K，Q，J和10。
	// 平手牌：公牌开出同花大顺，则所有未盖牌的牌手平手平分筹码。
	// 平分,且平點數,比同花色手牌

	// 同花顺（Straight Flush）
	// 五张同花色的连续数字牌。同时有同花顺时，数字最大者为赢家。
	// 平手牌：公牌开出同花顺为最大时，则所有未盖牌的牌手平手平分筹码。
	// 平分,且平點數,比同花色手牌

	// 铁支（四条，Four of a kind）
	// 其中四张是相同数字的扑克牌，第五张是剩下牌组中最大的一张牌。若有一家以上持有四条（公牌开出四条），则比较第五张牌（起脚牌），最大者为赢家。
	// 平手牌：公牌开出四条时，最后一张杂牌（或称为kicker、次大牌、踢脚牌，一副牌型组合中剩下来没有用作凑牌型的牌，用于牌型相同时比大小）数字也相同时。
	// 平分,且平點數,比手牌

	// 葫芦（夫佬或满堂红，Full house）
	// 由三张相同数字及任何两张其他相同数字的扑克牌组成，如果同时有多人拿到葫芦，三张相同数字中数字较大者为赢家。如果使用多副牌且三张牌都一样，则再比两张牌中数字较大者赢家。
	// 平手牌：五张牌数字都一样，则平分彩池。

	// 同花（Flush）
	// 此牌由五张不按顺序但相同花色的扑克牌组成，如果不只一人有此牌组，则牌面数字最大的人赢得该局，如果最大数字相同，则由第二、第三、第四或者第五张牌来决定胜负。
	// 平手牌：公牌的同花就是最大的同花牌型时，平分彩池。
	// 比同花色大小

	// 顺子（Straight）
	// 此牌由五张连续数字扑克牌组成，如果不只一人有此牌组，则五张牌中数字最大的赢得此局，10-J-Q-K-A为最大的顺子，A-2-3-4-5为最小的顺子。
	// 平手牌：如果五张牌数字都相同，平分彩池。

	// 三条（Three of a kind）
	// 由三张相同数字和两张不同数字的扑克牌组成，如果不只一人有此牌组，则三张牌中数字者最大赢得该局。如果使用多副牌且三张牌数字大小相同，则比较不同点数的两张牌中数字较大者，若相同时再比第五张，数字大的人赢。
	// 平手牌：如果五张牌数字都相同，则平分彩池。
	//比後兩張

	// 两对（Two pair）
	// 两对数字相同但两两不同的扑克和一张杂牌组成，共五张牌。
	// 平手牌：如果不只一人持有此牌型，持有数字比较大的对子者为赢家，若较大数字对子相同，则比较小对子的数字，如果两对对子数字都相同，那么第五张牌（kicker）数字较大者赢。如果连第五张牌数字也相同，则平分彩池。
	//比最後一張

	// 对子（Pair）
	// 由两张相同数字的扑克牌和另三张无法组成牌型的杂牌组成。
	// 平手牌：如果不只一人持有此牌型，则持有较大数字对子者为赢家，如果对牌数字相同，则依序比较剩下的三张牌，数字最大者为赢家，如果五张牌都一样，则平分彩池。
	// 比最後三張

	// 乌龙（高牌或散牌, High card，No-pair，Zilch）
	// 无法组成以上任一牌型的杂牌。
	// 平手牌：如果不只一人抓到此牌，则比较数字最大者，如果数字最大的相同，则依序比较第二、第三、第四和第五大的，如果五张牌都相同，则平分彩池。
	// 比所有牌
	return false
}

func setLog(proCard card, saveLog chan card) {
	proCard.getUserType()
	saveLog <- proCard
}

func (c *card) getUserType() {
	c.repeat()
	c.straight()
}

func (c *card) repeat() {
	statistics := make(map[int]int)
	for _, v := range c.input {
		statistics[getLowestDigit(v)]++
	}

	appearAgain := ""
	savePoint := []int{}
	for point, count := range statistics {
		if count > 0 {
			if count == 4 {
				c.cardType = 4
				c.cardPoint = point
				return
			} else if count == 3 {
				savePoint = append(savePoint, point)
				appearAgain += "3"
			} else if count == 2 {
				c.cardPoint = point
				savePoint = append(savePoint, point)
				appearAgain += "2"
			}
		}
	}

	switch appearAgain {
	case "32", "23":
		if savePoint[0] > savePoint[1] {
			c.cardPoint = savePoint[0]
			c.cardSecPoint = savePoint[1]
		} else {
			c.cardPoint = savePoint[1]
			c.cardSecPoint = savePoint[0]
		}

		c.cardType = 4

	case "3":
		c.cardType = 7
	case "22":
		if savePoint[0] > savePoint[1] {
			c.cardPoint = savePoint[0]
			c.cardSecPoint = savePoint[1]
		} else {
			c.cardPoint = savePoint[1]
			c.cardSecPoint = savePoint[0]
		}
		c.cardType = 8
	case "2":
		c.cardType = 9
	default:
		c.cardType = 10
	}
}

func (c *card) straight() {
	isStraight := true
	c.cardType = 10
	sortInput := []int{}
	for _, v := range c.input {
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

	c.flush()

	if isStraight && c.isFlish {
		c.calculatorStraight()
	} else if c.isFlish {
		c.cardType = 5
	} else if isStraight {
		c.cardPoint = recodeStraight[len(recodeStraight)-1]
		c.cardType = 6
	} else {
		c.cardType = 10
	}

}

func (c *card) flush() {
	flushMap := make(map[int]int)
	for _, v := range c.input {
		flushMap[getHighestDigit(v)]++
	}

	for _, v := range flushMap {
		if v >= 5 {
			c.isFlish = true
		}
	}

}

func (c *card) calculatorStraight() {
	recodeStraight := []int{c.input[0]}
	for i := 1; i < len(c.input); i++ {
		if c.input[i] != c.input[i-1]+1 {
			if getLowestDigit(c.input[i]) == 14 && getLowestDigit(c.input[i-1]) == 5 {
				c.input[i] = 1
				recodeStraight = append(recodeStraight, c.input[i])
				continue
			}

			if len(c.input)-i >= 5 {
				recodeStraight = []int{c.input[i]}
				continue
			}

			if len(recodeStraight) >= 5 {
				break
			}
			c.cardType = 5
			return
		}
		recodeStraight = append(recodeStraight, c.input[i])
	}

	sort.Ints(recodeStraight)
	c.cardPoint = getLowestDigit(recodeStraight[len(recodeStraight)-1])
	if c.cardPoint == 14 {
		c.cardType = 1
	} else {
		c.cardType = 2
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

// remove 刪除指定牌
func remove(slice []int, elems []int) []int {
	result := []int{}
	for _, v := range slice {
		if !contains(elems, v) {
			result = append(result, v)
		}
	}
	return result
}

func contains(slice []int, elem int) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
