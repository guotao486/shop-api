/*
 * @Author: GG
 * @Date: 2022-10-20 16:05:44
 * @LastEditTime: 2022-10-20 16:12:24
 * @LastEditors: GG
 * @Description:csv 读取工具类
 * @FilePath: \shop-api\utils\csv_helper\read_csv.go
 *
 */
package csv_helper

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
)

/**
 * @name: 从给定的FileHeader读信息，返回二维数组
 * @description:
 * @param {*multipart.FileHeader} fileHeader
 * @return {*}
 */
func ReadCsv(fileHeader *multipart.FileHeader) ([][]string, error) {
	// 打开文件
	f, err := fileHeader.Open()

	if err != nil {
		log.Print(err)
	}

	// 结尾关闭文件
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Print(err)
		}
	}(f)

	// 读取对象
	reader := csv.NewReader(f)
	// 读取全部
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("err")
		return nil, err
	}

	var result [][]string

	for _, line := range lines[1:] {
		data := []string{line[0], line[1]}
		result = append(result, data)
	}
	return result, nil
}
