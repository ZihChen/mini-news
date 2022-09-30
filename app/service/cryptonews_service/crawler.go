package cryptonews_service

import (
	"fmt"
	"github.com/gocolly/colly"
	"mini-news/app/global/errorcode"
	"mini-news/app/global/helper"
	"strings"
)

var c *colly.Collector

func init() {
	c = colly.NewCollector(
		colly.AllowedDomains(ABMEDIA_DOMAIN, BLOCKTEMPO_DOMAIN),
		colly.UserAgent(USER_AGENT),
	)
}

func (s *service) GetAbmediaArticles() (articles []CryptoArticle, goErr errorcode.Error) {
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div[class='loop-post']", func(e *colly.HTMLElement) {

		e.ForEach("article", func(i int, el *colly.HTMLElement) {

			src := el.ChildAttr("div[class='description'] > h3[class='title'] > a", "href")
			articles = append(articles, CryptoArticle{
				Name:   parsArticleName(src),
				Site:   ABMEDIA,
				Source: src,
				Title:  el.ChildText("div[class='description'] > h3"),
				Image: func() string {
					img := el.ChildAttr("figure > a > img", "data-lazy-srcset")
					parse := strings.Split(img, "?")
					return parse[0]
				}(),
				Date: func() string {
					datetime := el.ChildAttr("div[class='description'] > div > time", "datetime")
					formatStr := strings.Split(datetime, "T")
					return formatStr[0]
				}(),
			})

		})
	})

	c.OnError(func(response *colly.Response, err error) {
		goErr = helper.ErrorHandle(errorcode.ErrorService, errorcode.GetAbmediaNewsError, err.Error())
	})

	if goErr != nil {
		return
	}

	if err := c.Visit("https://" + ABMEDIA_DOMAIN + "/blog"); err != nil {
		goErr = helper.ErrorHandle(errorcode.ErrorService, errorcode.VisitAbmediaNewsError, err.Error())
		return
	}

	return
}

func (s *service) GetBlockTempoArticles() (articles []CryptoArticle, goErr errorcode.Error) {
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div[class='jeg_block_container']", func(e *colly.HTMLElement) {

		e.ForEach("article", func(i int, el *colly.HTMLElement) {

			src := el.ChildAttr("div[class='jeg_postblock_content'] > h3[class='jeg_post_title'] > a", "href")
			articles = append(articles, CryptoArticle{
				Name:   parsArticleName(src),
				Site:   BLOCKTEMPO,
				Source: src,
				Title:  el.ChildText("div[class='jeg_postblock_content'] > h3[class='jeg_post_title'] > a"),
				Image:  el.ChildAttr("div[class='jeg_thumb'] > a > div[class='thumbnail-container animate-lazy  size-715 '] > img", "data-src"),
				Date:   el.ChildText("div[class='jeg_postblock_content'] > div[class='jeg_post_meta'] > div[class='jeg_meta_date'] > a"),
			})
		})
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	if err := c.Visit("https://" + BLOCKTEMPO_DOMAIN + "/category/cryptocurrency-market"); err != nil {

	}

	return
}

func (s *service) GetBlockCastArticles() (articles []CryptoArticle, goErr errorcode.Error) {
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div[data-id='eaf3a42']", func(e *colly.HTMLElement) {

		e.ForEach("article", func(i int, el *colly.HTMLElement) {

			src := el.ChildAttr("div[class='jeg_postblock_content'] > h3[class='jeg_post_title'] > a", "href")
			articles = append(articles, CryptoArticle{
				Name:   parsArticleName(src),
				Site:   BLOCKCAST,
				Source: src,
				Title:  el.ChildText("div[class='jeg_postblock_content'] > h3[class='jeg_post_title'] > a"),
				Image:  el.ChildAttr("div[class='jeg_thumb'] > a > div[class='thumbnail-container animate-lazy  size-500 '] > img", "src"),
				Date:   el.ChildText("div[class='jeg_postblock_content'] > div[class='jeg_post_meta'] > div[class='jeg_meta_date'] > a"),
			})
		})
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	if err := c.Visit(BLOCKCAST_DOMAIN); err != nil {

	}

	return
}

func parsArticleName(src string) (articleName string) {
	str := strings.Split(strings.TrimRight(src, "/"), "/")

	return str[len(str)-1]
}
