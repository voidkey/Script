package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	// 打开CSV文件
	name := "bot"
	inputFile, err := os.Open(name + ".csv")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	// 创建输出文件
	outputFile, err := os.Create(name + "_output.csv")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// 遍历每一行记录
	for i, record := range records {
		// 读取mongoDB中的ISODate类型的时间字段,index是该字段的位置
		index := 0
		isoDateStr := record[index]
		mysqlDateStr := isoDateToMySQLDate(isoDateStr)

		// 更新记录中的日期
		record[index] = mysqlDateStr

		// 更新正确格式到输出文件
		if err := writer.Write(record); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write line %d: %v\n", i+1, err)
			continue // 跳过不能写入的行
		}
	}
}

func isoDateToMySQLDate(isoDateStr string) string {
	// 解析ISODate字符串为Go的time.Time对象
	isoDate, err := time.Parse(time.RFC3339, isoDateStr)
	if err != nil {
		return ""
	}

	// 转换为MySQL的DATETIME格式的字符串
	mysqlDateStr := isoDate.Format("2006-01-02 15:04:05")
	return mysqlDateStr
}
