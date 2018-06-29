package config

import (
	"math/rand"
	"time"
)

// ItemRuleConfig contains information about novel's rules
type ItemRuleConfig struct {
	NovelName              string
	NovelUrl               string
	NovelType              string
	NovelAuthor            string
	NovelCover             string
	NovelAbstract          string
	NovelLatestChapterName string
	NovelLatestChapterUrl  string
}

// NovelRule contains information about novle's source
type NovelRule struct {
	Name       string
	HomeUrl    string
	SearchUrl  string
	TargetItem string
	ItemRule   ItemRuleConfig
}

var (
	UserAgents = []string{
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7)",
		"Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 5.1; Trident/5.0)",
		"Mozilla/5.0 (X11; Linux x86_64; rv:2.2a1pre) Gecko/20100101 Firefox/4.2a1pre",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b11pre) Gecko/20110131 Firefox/4.0b11pre",
		"Mozilla/5.0 (X11; U; Linux i686; ru-RU; rv:1.9.2a1pre) Gecko/20090405 Ubuntu/9.04 (jaunty) Firefox/3.6a1pre",
		"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.8) Gecko/20100723 SUSE/3.6.8-0.1.1 Firefox/3.6.8",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; pt-PT; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; it; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 ( .NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 (.NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; ru; rv:1.9.2.4) Gecko/20100513 Firefox/3.6.4",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; ja; rv:1.9.2.4) Gecko/20100611 Firefox/3.6.4 GTB7.1",
	}
	// Common rules
	// 关于笔趣阁的规则
	ItemRule01 = ItemRuleConfig{
		NovelName:              "div.bookinfo h4.bookname a",
		NovelUrl:               "div.bookinfo h4.bookname a",
		NovelType:              "div.bookinfo div.cat",
		NovelAuthor:            "div.bookinfo div.author",
		NovelCover:             "div.bookimg a img",
		NovelAbstract:          "div.bookinfo p",
		NovelLatestChapterName: "div.bookinfo div.update a",
		NovelLatestChapterUrl:  "div.bookinfo div.update a",
	}
	// 百度第三方网站内容检索的规则
	ItemRule02 = ItemRuleConfig{
		NovelName:              "div.result-game-item-detail a.result-game-item-title-link",
		NovelUrl:               "div.result-game-item-detail a.result-game-item-title-link",
		NovelType:              "div.result-game-item-info p:nth-child(2)",
		NovelAuthor:            "div.result-game-item-info p:nth-child(1)",
		NovelCover:             "img.result-game-item-pic-link-img",
		NovelAbstract:          "p.result-game-item-desc",
		NovelLatestChapterName: "p.result-game-item-info-tag a.result-game-item-info-tag-item",
		NovelLatestChapterUrl:  "p.result-game-item-info-tag a.result-game-item-info-tag-item",
	}

	NovelsRulesMap = map[string]NovelRule{
		"10": NovelRule{
			Name:       "笔趣阁01",
			HomeUrl:    "https://www.bqg99.cc/",
			SearchUrl:  "https://www.bqg99.cc/s.php?q=",
			TargetItem: ".bookbox",
			ItemRule:   ItemRule01,
		},
		"11": NovelRule{
			Name:       "笔趣阁02",
			HomeUrl:    "http://www.cdzdgw.com/",
			SearchUrl:  "http://www.cdzdgw.com/s.php?q=",
			TargetItem: ".bookbox",
			ItemRule:   ItemRule01,
		},
		"12": NovelRule{
			Name:       "笔趣阁03",
			HomeUrl:    "http://www.biqugex.com/",
			SearchUrl:  "http://www.biqugex.com/s.php?q=",
			TargetItem: ".bookbox",
			ItemRule:   ItemRule01,
		},
		"20": NovelRule{
			Name:       "新笔趣阁01",
			HomeUrl:    "http://www.biqugetv.com/",
			SearchUrl:  "http://zhannei.baidu.com/cse/search?click=1&s=16765504158186272814&q=",
			TargetItem: "div.result-list div.result-item",
			ItemRule:   ItemRule02,
		},
		"30": NovelRule{
			Name:       "笔下文学01",
			HomeUrl:    "http://www.xbxwx.net/",
			SearchUrl:  "http://so.xbxwx.net/cse/search?click=1&entry=1&s=10874778206555383279&q=",
			TargetItem: "div.result-list div.result-item",
			ItemRule:   ItemRule02,
		},
	}
)

// GetUserAgent returns a random user agent
func GetUserAgent() string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(UserAgents)
	return UserAgents[n]
}
