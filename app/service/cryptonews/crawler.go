package cryptonews

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

var c *colly.Collector

func init() {
	c = colly.NewCollector(
		colly.AllowedDomains(ABMEDIA_DOMAIN, BLOCKTEMPO_DOMAIN),
		colly.UserAgent(USER_AGENT),
	)
}

func (s *service) GetAbmediaArticles() (articles []CryptoArticle) {
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div[class='loop-post']", func(e *colly.HTMLElement) {

		e.ForEach("article", func(i int, el *colly.HTMLElement) {

			src := el.ChildAttr("div[class='description'] > h3[class='title'] > a", "href")
			articles = append(articles, CryptoArticle{
				ArticleName: parsArticleName(src),
				Type:        ABMEDIA,
				Source:      src,
				Title:       el.ChildText("div[class='description'] > h3"),
				Image: func() string {
					img := el.ChildAttr("figure > a > img", "data-lazy-srcset")
					parse := strings.Split(img, "?")
					return parse[0]
				}(),
				Time: el.ChildAttr("div[class='description'] > div > time", "datetime"),
			})

		})
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	if err := c.Visit("https://" + ABMEDIA_DOMAIN + "/blog"); err != nil {

	}

	return
}

func (s *service) GetBlockTempoArticles() (articles []CryptoArticle) {
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div[class='jeg_block_container']", func(e *colly.HTMLElement) {

		e.ForEach("article", func(i int, el *colly.HTMLElement) {

			src := el.ChildAttr("div[class='jeg_postblock_content'] > h3[class='jeg_post_title'] > a", "href")
			articles = append(articles, CryptoArticle{
				ArticleName: parsArticleName(src),
				Type:        BLOCKTEMPO,
				Source:      src,
				Title:       el.ChildText("div[class='jeg_postblock_content'] > h3[class='jeg_post_title'] > a"),
				Image:       el.ChildAttr("div[class='jeg_thumb'] > a > div[class='thumbnail-container animate-lazy  size-715 '] > img", "data-src"),
				Time:        el.ChildText("div[class='jeg_postblock_content'] > div[class='jeg_post_meta'] > div[class='jeg_meta_date'] > a"),
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

func (s *service) GetBlockCastArticles() (articles []CryptoArticle) {
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div[data-id='eaf3a42']", func(e *colly.HTMLElement) {

		e.ForEach("article", func(i int, el *colly.HTMLElement) {

			src := el.ChildAttr("div[class='jeg_postblock_content'] > h3[class='jeg_post_title'] > a", "href")
			articles = append(articles, CryptoArticle{
				ArticleName: parsArticleName(src),
				Type:        BLOCKCAST,
				Source:      src,
				Title:       el.ChildText("div[class='jeg_postblock_content'] > h3[class='jeg_post_title'] > a"),
				Image:       el.ChildAttr("div[class='jeg_thumb'] > a > div[class='thumbnail-container animate-lazy  size-500 '] > img", "src"),
				Time:        el.ChildText("div[class='jeg_postblock_content'] > div[class='jeg_post_meta'] > div[class='jeg_meta_date'] > a"),
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
