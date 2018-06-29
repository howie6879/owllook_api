package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/howie6879/owllook_api/config"
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
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		raw_html := DetectBody(body)
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
func RequestURL(url string) (*http.Response, error) {
	tr := &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", config.GetUserAgent())
	response, err := client.Do(req)
	return response, err
}
