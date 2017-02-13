package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// 武汉通API响应结构
type whtong struct {
	Status int  `json:"status"`
	Data   data `json:"data"`
}

type data struct {
	ID     string   `json:"id"`
	YuE    string   `json:"yue"`
	Record []record `json:"record"`
}

type record struct {
	Time   string `json:"time"`
	Type   string `json:"type"`
	Amount string `json:"amount"`
}

func getId() string {
	var id uint64
	fmt.Printf("请输入需要查询的武汉通卡ID: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		fmt.Println("您输入的ID有误")
	}
	fmt.Println("您输入的ID为:", id)
	return strconv.FormatUint(id, 10)
}

func getInfo(id string) (*whtong, error) {
	url := "https://wht.whutech.com/data?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("似乎通讯不畅,请稍后尝试")
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("数据读取出错")
		return nil, err
	}

	wht := &whtong{}
	if err := json.Unmarshal(body, wht); err != nil {
		fmt.Println("数据解析出错")
		return nil, err
	}
	return wht, nil
}

func printResult(wht *whtong) {
	fmt.Println("武汉通查询结果:")
	fmt.Println("=========================================")
	fmt.Printf("%s: %s\n", "卡号", wht.Data.ID)
	fmt.Printf("%s: %s %s\n", "余额", wht.Data.YuE, "元")
	fmt.Printf("%s: \n", "记录")
	for _, v := range wht.Data.Record {
		fmt.Println(v)
	}
	fmt.Println("=========================================")
}

func main() {
AGAIN:
	id := getId()
	wht, err := getInfo(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if wht.Status == 200 {
		printResult(wht)
	} else {
		fmt.Println("您输入的武汉通ID有误!")
	}
	fmt.Print("继续查询请输入Y，其他输入将退出程序: ")
	isAgain := ""
	fmt.Scanf("%s", &isAgain)
	if strings.EqualFold(isAgain, "Y") || strings.EqualFold(isAgain, "y") {
		goto AGAIN
	} else {
		return
	}
}
