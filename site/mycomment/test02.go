package main

import (
	"ToolExcelize"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

var max *string = flag.String("max", "29", "movieid")
var min *string = flag.String("min", "1", "movieid")
var title *string = flag.String("title", "SIGN", "movieid")

func ExampleScrape(hhh int, hhhh int, hhhhh string) {
	xlsx, err := excelize.OpenFile("abc.xlsx")
	if err != nil {
		fmt.Println("当前文件夹下 找不到 abc.xlsx")
		os.Exit(1)
	}
	fmt.Println("=============================开始采集")
	for i := hhhh; i <= hhh; i++ {
		AA := "A" + strconv.Itoa(i)
		newstitle := hhhhh + strconv.Itoa(i)
		gqygqy := string(xlsx.GetCellValue("Sheet1", AA)) //gqygqy := "http://auto.huanqiu.com/globalnews/2017-07/11036229.html"
		//
		re, _ := regexp.Compile("^http")
		one := re.Find([]byte(gqygqy))
		if one == nil {

		} else {
			time.Sleep(12 * time.Millisecond)
			doc, err := goquery.NewDocument(gqygqy)
			if err != nil {
				fmt.Println("getdoc 异常", err)
				continue
			}
			doc.Find("body").Each(func(i int, s *goquery.Selection) {
				title := s.Find("div.article div.article").Text()
				if title == "" {
					title = s.Find("article.article").Text()
				}
				if title == "" {
					title = s.Find("div.article").Text()
				}
				if title == "" {
					title = s.Find("div.article-cont").Text()
				}
				if title == "" {
					title = s.Find("div.ina_content").Text()
				}
				if title == "" {
					title = s.Find("div.article-body").Text()
				}
				if title == "" {
					title = s.Find("div.Article").Text()
				}
				if title == "" {
					title = s.Find("div.arl-c-txt").Text()
				}
				if title == "" {
					title = s.Find("div#article_text").Text()
				}
				if title == "" {
					title = s.Find("div#contentText").Text()
				}
				if title == "" {
					title = s.Find("div.rich_media_content ").Text()
				}
				if title == "" {
					title = s.Find("div#con_wrap div#content").Text()
				}
				if title == "" {
					title = s.Find("div.article-mod div.at-cnt-main").Text()
				}
				if title == "" {
					title = s.Find("div#ArticleContent").Text()
				}
				if title == "" {
					title = s.Find("div.article-content example").Text()
				}
				if title == "" {
					title = s.Find("div.article-content").Text()
				}
				if title == "" {
					title = s.Find("div#news_content p").Text()
				}
				if title == "" {
					title = s.Find("div.news_con_text").Text()
				}

				if title == "" {
					title = s.Find("div.qq_article").Text()
				}
				if title == "" {
					title = s.Find("div#artibody").Text()
				}
				if title == "" {
					title = s.Find("div.news_body div.news_main").Text()
				}

				if title == "" {
					title = s.Find("div.inforcont div.dealermain").Text()
				}
				if title == "" {
					title = s.Find("div.ds_page_content").Text()
				}
				if title == "" {
					title = s.Find("div#chan_newsDetail").Text()
				}
				if title == "" {
					title = s.Find("div.art_text").Text()
				}
				if title == "" {
					title = s.Find("div.post_body div.post_text").Text()
				}
				if title == "" {
					title = s.Find("div.content_detail_left").Text()
				}
				if title == "" {
					title = s.Find("div#dvContent").Text()
				}
				if title == "" {
					title = s.Find("div.news_container").Text()
				}
				if title == "" {
					title = s.Find("div.content-box").Text()
				}

				if title == "" {
					title = s.Find("div.content").Text()
				}
				if title == "" {
					title = s.Find("div#content").Text()
				}
				if title == "" {
					title = s.Find("div#ctrlfscont").Text()
				}
				if title == "" {
					title = s.Find("div.main_fl").Text()
				}
				if title == "" {
					title = s.Find("div.mainBox").Text()
				}
				if title == "" {
					title = s.Find("div").Text()
				}
				if title == "" {
					title = s.Find("p").Text()
				}
				re, _ := regexp.Compile("auto.163.com|dealer.autohome.com.cn|hea.e23.cn|163.com")
				one := re.Find([]byte(gqygqy))
				if one != nil {
					title = mahonia.NewDecoder("gbk").ConvertString(string(title))
				}
				postcontent := fmt.Sprintf(title, ".")
				fmt.Println("=============" + AA + "采集ing" + " =============")
				//开始保存
				if title != "" {
					file, createErr := os.Create(newstitle + ".txt")
					if createErr != nil {
						log.Fatal("创建文件失败", createErr)
					}

					_, writeErr := file.Write([]byte(postcontent))
					if writeErr != nil {
						log.Fatal(writeErr)
					}
					closeErr := file.Close()
					if closeErr != nil {
						log.Fatal(closeErr)
					}
				}
				//保存结束
			})
		}
	}
	fmt.Println("=============================采集已结束")
	//	reg := regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|]`)
	//	src =reg.ReplaceAllString(src,"")
}

func main() {
	flag.Parse()
	hhh, _ := strconv.Atoi(*max)
	hhhh, _ := strconv.Atoi(*min)
	hhhhh := *title
	ExampleScrape(hhh, hhhh, hhhhh)

}
