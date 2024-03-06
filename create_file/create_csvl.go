package create_file

import (
	"encoding/csv"
	"log"
	"os"
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


func CreateCsv(data [][]string) {
	// 資料範例

	// 建立一個新的CSV檔案
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("無法建立檔案:", err)
	}
	defer file.Close()

	// 建立CSV寫入器
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 寫入資料
	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			log.Fatal("寫入CSV時發生錯誤:", err)
		}
	}

	log.Println("資料成功寫入CSV檔案")
}
