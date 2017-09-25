package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func getTotalNum(a int) int {
	baseurll := "http://you.163.com/xhr/comment/listByItemByTag.json?&itemId=" + strconv.Itoa(a) + "&tag=%E5%85%A8%E9%83%A8&size=20&page=1&orderBy=1"
	resp, err := http.Get(baseurll)
	if err != nil {

		fmt.Println("=======999====")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json, err := simplejson.NewJson(body)
	//fmt.Println(json)
	totalnum := json.Get("data").Get("pagination").Get("totalPage")
	totalnumhh := fmt.Sprintf("%s", totalnum)
	re, _ := regexp.Compile("(\\d+)")
	one := re.Find([]byte(totalnumhh))
	total, _ := strconv.Atoi(string(one))
	if total > 1001 {
		total = 1000
	}
	return total
}

var id *string = flag.String("id", "1333007", "productId")
var initdate *string = flag.String("initdate", "1999/09/09", "initdate")

func main() {
	flag.Parse()
	movieid, _ := strconv.Atoi(*id)
	fmt.Println("===========you.163.com productId(cic):", movieid)
	timeLayout := "2006/01/02"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, *initdate, loc)
	sr := theTime.Unix()
	//fmt.Println("===========you.163.com productId(cic):", sr)
	maxnum := getTotalNum(movieid)
	if maxnum > 0 {
		t := time.Now().Unix()
		name := "you.163.com-" + *id + "-comment-" + strconv.FormatInt(t, 10) + ".csv"
		fmt.Println(name)
		csvFile, err := os.Create(name)
		if err != nil {
			panic("创建文件失败")
		}
		defer csvFile.Close()
		csvFile.WriteString("\xEF\xBB\xBF")
		writer := csv.NewWriter(csvFile)

		for page := 1; page <= maxnum; page++ {
			stringpage := strconv.Itoa(page)
			baseurl := "http://you.163.com/xhr/comment/listByItemByTag.json?&itemId=" + *id + "&tag=%E5%85%A8%E9%83%A8&size=20&page=" + stringpage + "&orderBy=1"
			time.Sleep(2 * time.Second)
			fmt.Println("第", stringpage, "页")
			resp, err := http.Get(baseurl)
			if err != nil {

				fmt.Println("=======999====")
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			json, err := simplejson.NewJson(body)
			data, err := json.Get("data").Get("result").Array()
			//fmt.Println(data)
			if page == 1 {
				lineo := []string{"poster", "createTime", "content", "star", "memberLevel", "Info"}
				err := writer.Write(lineo)
				if err != nil {
					panic("panic GQY")
				}
			}
			for _, v := range data {
				vv := v.(map[string]interface{})
				frontUserName := vv["frontUserName"].(string)
				content := vv["content"].(string)
				tm := fmt.Sprintf("%s", vv["createTime"])
				rs := []rune(tm)
				tmqq := string(rs[0:10])
				tmm, _ := strconv.Atoi(tmqq)
				if int64(tmm) < sr {
					time.Sleep(2 * time.Second)
					fmt.Println("===========采集结束==")
					writer.Flush()
					return
				}
				tmmm := time.Unix(int64(tmm), 0)
				posttime := tmmm.Format("2006-01-02 15:04:05")

				star := fmt.Sprintf("%s", vv["star"])
				skuInfo := fmt.Sprintf("%s", vv["skuInfo"])
				//				approvenum := fmt.Sprintf("%s", vv["approve"])
				memberLevel := fmt.Sprintf("%s", vv["memberLevel"])
				line := []string{frontUserName, posttime, content, star, memberLevel, skuInfo} //, replynum, approvenum, score}
				fmt.Println(line)

				err := writer.Write(line)
				if err != nil {
					panic("panic GQY")
				}
			}
			writer.Flush()
		}
	}
	time.Sleep(2 * time.Second)
	fmt.Println("===========采集结束")
}
