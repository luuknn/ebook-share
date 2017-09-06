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
	baseurll := "http://m.maoyan.com/comments.json?movieid=%d&limit=10&offset=0"
	resp, err := http.Get(fmt.Sprintf(baseurll, a))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json, err := simplejson.NewJson(body)
	totalnum := json.Get("data").Get("CommentResponseModel").Get("total")
	totalnumhh := fmt.Sprintf("%s", totalnum)
	re, _ := regexp.Compile("(\\d+)")
	one := re.Find([]byte(totalnumhh))
	total, _ := strconv.Atoi(string(one))
	if total > 1001 {
		total = 1000
	}
	return total
}

var id *string = flag.String("id", "344264", "movieid")

func main() {
	flag.Parse()
	movieid, _ := strconv.Atoi(*id)
	fmt.Println("===========猫眼ID号(cic):", movieid)
	//movieid := 344264
	maxnum := getTotalNum(movieid)
	t := time.Now().Unix()
	name := "maoyan-post-" + strconv.FormatInt(t, 10) + ".csv"
	fmt.Println(name)
	csvFile, err := os.Create(name)
	if err != nil {
		panic("创建文件失败")
	}
	defer csvFile.Close()
	csvFile.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(csvFile)

	for page := -1; 10*page < maxnum; page++ {
		baseurl := "http://m.maoyan.com/comments.json?movieid=%d&limit=10&offset=%d"
		baseurls := "http://m.maoyan.com/movie/%d.json"
		time.Sleep(1 * time.Second)
		resp, err := http.Get(fmt.Sprintf(baseurl, movieid, page*10))
		if page < 0 {
			time.Sleep(1 * time.Second)
			resp, err = http.Get(fmt.Sprintf(baseurls, movieid))
		}
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		json, err := simplejson.NewJson(body)
		data, err := json.Get("data").Get("CommentResponseModel").Get("cmts").Array()
		if page < 0 {
			data, err = json.Get("data").Get("CommentResponseModel").Get("hcmts").Array()
		}

		//fmt.Println(data)
		for _, v := range data {
			vv := v.(map[string]interface{})
			nick := vv["nickName"].(string)
			cont := vv["content"].(string)
			posttime := vv["time"].(string)
			replynum := fmt.Sprintf("%s", vv["reply"])
			approvenum := fmt.Sprintf("%s", vv["approve"])
			score := fmt.Sprintf("%s", vv["score"])
			line := []string{nick, posttime, cont, replynum, approvenum, score}
			fmt.Println(line)
			err := writer.Write(line)
			if err != nil {
				panic("panic GQY")
			}
		}
		writer.Flush()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("===========采集结束")
}
