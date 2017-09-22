package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

func getTotalNum(a int) int {
	baseurll := "https://www.kaola.com/commentAjax/comment_list.html"

	resp, err := http.PostForm(baseurll, url.Values{"goodsId": {*id}, "pageNo": {"1"}, "pageSize": {"20"}})
	if err != nil {

		fmt.Println("=======999====")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json, err := simplejson.NewJson(body)
	totalnum := json.Get("commentPage").Get("totalPage")
	totalnumhh := fmt.Sprintf("%s", totalnum)
	re, _ := regexp.Compile("(\\d+)")
	one := re.Find([]byte(totalnumhh))
	total, _ := strconv.Atoi(string(one))
	if total > 1001 {
		total = 1000
	}
	return total
}

var id *string = flag.String("id", "1608638", "movieid")

func main() {
	flag.Parse()
	movieid, _ := strconv.Atoi(*id)
	fmt.Println("===========kaola productId(cic):", movieid)
	maxnum := getTotalNum(movieid)
	if maxnum > 0 {
		t := time.Now().Unix()
		name := "kaola-" + *id + "-comment-" + strconv.FormatInt(t, 10) + ".csv"
		fmt.Println(name)
		csvFile, err := os.Create(name)
		if err != nil {
			panic("创建文件失败")
		}
		defer csvFile.Close()
		csvFile.WriteString("\xEF\xBB\xBF")
		writer := csv.NewWriter(csvFile)

		for page := 1; page <= maxnum; page++ {
			baseurl := "https://www.kaola.com/commentAjax/comment_list.html"
			time.Sleep(1 * time.Second)
			stringpage := strconv.Itoa(page)
			resp, err := http.PostForm(baseurl, url.Values{"goodsId": {*id}, "pageNo": {stringpage}, "pageSize": {"20"}})
			if err != nil {

				fmt.Println("=======999====")
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			json, err := simplejson.NewJson(body)
			data, err := json.Get("commentPage").Get("result").Array()
			fmt.Println(data)
			for _, v := range data {
				vv := v.(map[string]interface{})
				nick := vv["nicknameKaola"].(string)
				cont := vv["commentContent"].(string)
				posttime := vv["createTime"].(string)
				alias := fmt.Sprintf("%s", vv["alias"])
				//				replynum := fmt.Sprintf("%s", vv["reply"])
				//				approvenum := fmt.Sprintf("%s", vv["approve"])
				score := fmt.Sprintf("%s", vv["commentPoint"])
				line := []string{nick, posttime, cont, score, alias} //, replynum, approvenum, score}
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
