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
)

var max *string = flag.String("max", "20", "movieid")

func ExampleScrape() {
	flag.Parse()
	hhh, _ := strconv.Atoi(*max)
	xlsx, err := excelize.OpenFile("abc.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("=============================开始采集")
	for i := 2; i <= hhh; i++ {
		AA := "A" + strconv.Itoa(i)
		gqygqy := string(xlsx.GetCellValue("Sheet1", AA)) //gqygqy := "http://auto.huanqiu.com/globalnews/2017-07/11036229.html"
		//
		re, _ := regexp.Compile("^http")
		one := re.Find([]byte(gqygqy))
		if one == nil {

		} else {
			//
			doc, err := goquery.NewDocument(gqygqy)
			if err != nil {
				log.Fatal(err)
			}
			doc.Find("body").Each(func(i int, s *goquery.Selection) {
				title := s.Find("div.article").Text()
				if title == "" {
					title = s.Find("article.article").Text()
				}
				if title == "" {
					title = s.Find("div.article-cont").Text()
				}
				if title == "" {
					title = s.Find("div.ina_content").Text()
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
					title = s.Find("div.content-box").Text()
				}
				if title == "" {
					title = s.Find("div.qq_article").Text()
				}
				if title == "" {
					title = s.Find("div.art_text").Text()
				}
				if title == "" {
					title = s.Find("div.content_detail_left").Text()
				}
				if title == "" {
					title = s.Find("div.content").Text()
				}
				if title == "" {
					title = s.Find("div#content").Text()
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
				utf8 := mahonia.NewDecoder("gbk").ConvertString(string(title))
				postcontent := fmt.Sprintf("URL:%s\n UTF-8 START:%s\n GBK START:%s\n", gqygqy, title, utf8)
				fmt.Println("=============" + AA + "采集ing" + " =============")
				//开始保存
				if title != "" {
					file, createErr := os.Create(AA + ".txt")
					if createErr != nil {
						log.Fatal(createErr)
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
}

func main() {
	ExampleScrape()
}
