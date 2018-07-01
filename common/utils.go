package common

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/howie6879/owllook_api/config"
	"github.com/levigross/grequests"
	"github.com/saintfish/chardet"
)

// DetectBody gbk convert to utf-8
func DetectBody(body []byte) string {
	var bodyString string
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(body)
	if err != nil {
		return string(body)
	}
	if strings.Contains(strings.ToLower(result.Charset), "utf") {
		bodyString = string(body)
	} else {
		bodyString = mahonia.NewDecoder("gbk").ConvertString(string(body))
	}
	return bodyString
}

// MakeAbsolute returns a absolute url
func MakeAbsolute(homeUrl string, currentUrl string) string {
	urlParse, _ := url.Parse(currentUrl)
	urlHost := urlParse.Host
	homeUrlParse, _ := url.Parse(homeUrl)
	if urlHost == "" {
		absoluteUrl := homeUrlParse.ResolveReference(urlParse)
		return absoluteUrl.String()
	}
	return urlParse.String()
}

// FetchHtml returns a raw html
func FetchHtml(name string, rule config.NovelRule) ([]map[string]string, error) {
	url := rule.SearchUrl + name
	response, err := RequestURL(url)
	var resultData []map[string]string
	if err != nil {
		log.Println("Request URL error", err)
		return resultData, err
	}
	if response.StatusCode == 200 {
		raw_html := DetectBody(response.Bytes())
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(raw_html))
		doc.Find(rule.TargetItem).Each(func(i int, s *goquery.Selection) {
			novelName := s.Find(rule.ItemRule.NovelName).Text()
			novelUrl, _ := s.Find(rule.ItemRule.NovelUrl).Attr("href")
			absoluteNovelUrl := MakeAbsolute(rule.HomeUrl, novelUrl)
			novelType := s.Find(rule.ItemRule.NovelType).Text()
			novelCover, _ := s.Find(rule.ItemRule.NovelCover).Attr("src")
			absoluteNovelCover := MakeAbsolute(rule.HomeUrl, novelCover)
			novelAuthor := s.Find(rule.ItemRule.NovelAuthor).Text()
			novelAbstract := s.Find(rule.ItemRule.NovelAbstract).Text()
			novelLatestChapterName := s.Find(rule.ItemRule.NovelLatestChapterUrl).Text()
			novelLatestChapterUrl, _ := s.Find(rule.ItemRule.NovelLatestChapterUrl).Attr("href")
			absoluteNovelLatestChapterUrl := MakeAbsolute(rule.HomeUrl, novelLatestChapterUrl)
			currentItem := make(map[string]string)
			currentItem["source_name"] = rule.Name
			currentItem["source_url"] = rule.HomeUrl
			currentItem["novel_name"] = strings.TrimSpace(novelName)
			currentItem["novel_url"] = absoluteNovelUrl
			currentItem["novel_type"] = strings.TrimSpace(novelType)
			currentItem["novel_cover"] = absoluteNovelCover
			currentItem["novel_author"] = strings.TrimSpace(novelAuthor)
			currentItem["novel_abstract"] = strings.TrimSpace(novelAbstract)
			currentItem["novel_latest_chapter_name"] = novelLatestChapterName
			currentItem["novel_latest_chapter_url"] = absoluteNovelLatestChapterUrl
			resultData = append(resultData, currentItem)
		})
	}
	return resultData, nil
}

// RequestURL returns a search result
func RequestURL(url string) (*grequests.Response, error) {
	ro := &grequests.RequestOptions{
		Headers: map[string]string{"User-Agent": config.GetUserAgent()},
	}
	resp, err := grequests.Get(url, ro)
	if err != nil {
		log.Println("Unable to make request: ", err)
	}

	// log.Println(resp.String())
	return resp, err
}
