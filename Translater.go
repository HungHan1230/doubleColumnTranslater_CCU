package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jinzhongmin/gtra"
)

func main() {
	// 開檔
	inputFile, Error := os.Open("paper.txt")

	fmt.Println("reading...")
	// 判斷是否開檔錯誤
	if Error != nil {
		fmt.Println("開檔錯誤!")
		return
	}
	// 離開時自動執行關檔
	defer inputFile.Close()
	// 宣告行計數器
	cnt := 1
	//
	inputReader := bufio.NewReader(inputFile)
	//^^^^^^^^^^^         ^^^^^^^   ^^^^^^^
	// 緩衝輸入物件        建立函數   來源:已開啟檔案

	// 用迴圈讀取檔案內容
	for {
		// 讀取字串直到遇到跳行符號
		//inputString, Error := inputReader.ReadString('\n')
		inputString, Error := inputReader.ReadString('.')
		// 若到檔尾時分發生  io.EOF 錯誤
		// 根據此錯誤 判斷是否離開
		if Error == io.EOF {
			fmt.Println("已讀取到檔尾!!")
			return
		}

		replaceStr := strings.Replace(inputString, "\n", " ", -1)

		t := gtra.NewTranslater()
		Error, gtraStr := t.Translate(replaceStr)
		if Error != nil {
			fmt.Println("Error!!")
			return
		}

		// 列印內容
		//fmt.Printf("第%2d行:%s\n", cnt, inputString)
		fmt.Println(replaceStr)
		fmt.Println(gtraStr + "\n")
		cnt++
	}
}
