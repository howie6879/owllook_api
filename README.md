## owllook小说接口

本项目提供优雅的小说接口，其他衍生项目如下：

- 公众号：[**粮草小说**](http://oe7yjec8x.bkt.clouddn.com/howie/2018-03-13-%E7%B2%AE%E8%8D%89%E5%B0%8F%E8%AF%B4.jpg-blog.howie)，有兴趣的话可以关注下
- 官网：[https://www.owllook.net](https://www.owllook.net)
- 监控工具：[owllook_gui](https://github.com/howie6879/owllook_gui)

### Overview

[owllook](https://github.com/howie6879/owllook)是一个基于其他搜索引擎构建的垂直小说搜索引擎，owllook目的是让阅读更简单、优雅，让每位读者都有舒适的阅读体验，有朋友有兴趣开发owllook的app端，于是此项目便诞生了，本项目提供小说的一系列接口，如检索、目录、章节内容、检查更新

#### Installation

`owllook_api`基于[gin](https://github.com/gin-gonic/gin)，提供了一系列基本的小说接口，使用：

``` shell
git clone https://github.com/howie6879/owllook_api
cd owllook_api

go get -u github.com/kardianos/govendor
govendor sync
go run main.go
```

#### API

小说资源说明：

本项目利用了互联网上的一些小说资源作为检索目标进行资源提取，定义如下，100以后的命名定位第三方检索：

| 名称 | 编号     |搜索类型		|
| :--- | -------- | ----------- |
| 10   | 笔趣阁   | 站内		|
| 100  | 新笔趣阁 |	百度第三方  |
| 110  | 笔下文学 | 百度第三方  |
| 120  | 顶点小说 | 百度第三方  |

**搜索小说：**

格式：/v1/novels/:name/:source

请求：

``` shell
# 请求资源为10下小说牧神记的检索结果
curl http://0.0.0.0:8080/v1/novels/牧神记/10
```

响应：

``` json
{
    "info": [
        {
            "novel_abstract": "大墟的祖训说，天黑，别出门。大墟残老村的老弱病残们从江边捡到了一个婴儿，取名秦牧，含辛茹苦将他养大。这一天夜幕降临，……",
            "novel_author": "作者：宅猪",
            "novel_cover": "https://www.bqg99.cc/bookimages/2640967.jpg",
            "novel_latest_chapter_name": "第七百九十二章 道一（月底求月票）",
            "novel_latest_chapter_url": "https://www.bqg99.cc/book/2639610/595030666.html",
            "novel_name": "牧神记",
            "novel_type": "分类：玄幻",
            "novel_url": "https://www.bqg99.cc/book/2639610/",
            "source_name": "笔趣阁01",
            "source_url": "https://www.bqg99.cc/"
        }
    ],
    "status": 1
}
```

