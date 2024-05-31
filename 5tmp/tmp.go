package main

import (
	"fmt"
	"sort"
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

func main() {
	fmt.Println(f([]int{0x10b, 0x10a}, []int{0x105, 0x104, 0x10c, 0x10d, 0x10e}))
}

func f(self, board []int) int {
	allCard := append(self, board...)

	card := calculatorCard{
		allCard:  allCard,
		isflush:  make(map[int]int),
		isRepeat: make(map[int]int),
	}

	for _, v := range allCard {
		card.pointArr = append(card.pointArr, getLowestDigit(v))
		card.flushArr = append(card.flushArr, getHighestDigit(v))
	}

	sort.Ints(card.pointArr)
	card.calculate()

	return card.cardType
}

type calculatorCard struct {
	allCard   []int
	pointArr  []int
	flushArr  []int
	cardType  int
	round     int

	isflush    map[int]int
	isRepeat   map[int]int
	isStraight []int
}

func (c *calculatorCard) calculate() {

	c.isflush[c.flushArr[c.round]]++

	if c.round == 0 {
		c.round++
		c.isStraight = append(c.isStraight, c.pointArr[0])
		c.calculate()
		return
	}

	if c.pointArr[c.round] == c.pointArr[c.round-1] {
		c.isRepeat[c.pointArr[c.round]]++
	} else if c.pointArr[c.round] == c.pointArr[c.round-1]+1 {
		c.isStraight = append(c.isStraight, c.pointArr[c.round])
	} else if c.pointArr[c.round] == c.pointArr[c.round-1]+9 && len(c.isStraight) == 4 {
		c.isStraight = append(c.isStraight, c.pointArr[c.round]-13)
	}

	if len(c.isStraight) >= 5 {
		c.isStraight = []int{}
	}

	if c.round == len(c.allCard)-1 {
		appearAgain := ""
		isFlush := false
		for _, v := range c.isRepeat {
			if v == 3 {
				c.cardType = 3
				return
			} else if v == 2 {
				appearAgain += "3"
			} else if v == 1 {
				appearAgain += "2"
			}
		}

		switch appearAgain {
		case "32", "23":
			c.cardType = 4
			return
		case "3":
			c.cardType = 7
		case "22":
			c.cardType = 8

		case "2":
			c.cardType = 9
		default:
			c.cardType = 10
		}

		for _, v := range c.isflush {
			if v >= 5 {
				isFlush = true
			}
		}

		if len(c.isStraight) >= 5 && isFlush {
			c.checkStraightFlush()
			return
		} else if isFlush {
			c.cardType = 5
		} else if len(c.isStraight) >= 5 {
			c.cardType = 2
		}
		return
	}
	c.round++
	c.calculate()
}

func (c *calculatorCard) checkStraightFlush() {
	sort.Ints(c.allCard)
	recodeStraight := []int{c.allCard[0]}
	fmt.Println(c.allCard)
	for i := 1; i < len(c.allCard); i++ {
		if c.allCard[i] != c.allCard[i-1]+1 {
			fmt.Println(getLowestDigit(c.allCard[i]), getLowestDigit(c.allCard[i-1]))
			if getLowestDigit(c.allCard[i]) == 14 && getLowestDigit(c.allCard[i-1]) == 5 {
				c.allCard[i] = 1
				recodeStraight = append(recodeStraight, c.allCard[i])
				continue
			}

			if len(c.allCard)-i >= 5 {
				recodeStraight = []int{c.allCard[i]}
				continue
			}

			if len(recodeStraight) >= 5 {
				break
			}

			c.cardType = 5
			return
		}
		recodeStraight = append(recodeStraight, c.allCard[i])
		fmt.Println(recodeStraight)
	}

	sort.Ints(recodeStraight)
	if getLowestDigit(recodeStraight[len(recodeStraight)-1]) == 14 {
		c.cardType = 1
		return
	} else {
		c.cardType = 2
		return
	}
}

func getLowestDigit(hexNumber int) int {
	return hexNumber & 0xF
}

func getHighestDigit(hexNumber int) int {
	firstDigit := hexNumber >> 8
	firstDigit = firstDigit & 0xF
	return firstDigit
}
