package colly

import (
	"bytes"
	"context"
	"fmt"
	"github.com/WooRho/rhtool/rhtool_core/rexcel"
	snowflack "github.com/WooRho/rhtool/rhtool_snowflack"
	"github.com/gocolly/colly"
	. "go_study/mysql"
	"go_study/structure"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func proxyAllPageInfo() {
	var (
		pageInfoList = make(PageInfoList, 0)
		buffer       = &bytes.Buffer{}
		randNumber   time.Duration
		FairTime     int
	)
	baseInfo := Init("a=51krykFTJ7qI2tOlUmVbedSVcZw7h04c; _popprepop=1; zone-cap-3576187=1%3B1711166608; token_QpUJAAAAAAAAGu98Hdz1l_lcSZ2rY60Ajjk9U1c=BQPyAAAAAAAACZUAApG6kSTWc7nWCL1mRgpNrrpwkRwqeKbvFUYUxMJzzG6MticixdK-mNKN8oI3e4qbBcpXk25D5hX3TUsfV7zgEfcC3bVFvsJw6GvzhX-J8tY7GKL5lT4TQEUUNGzj6f0pAHmySNM3HV9qxfWv4uM6Uu8fZ4k4G2C9_Uy50gKWc-_R6X1_9pHd0P-QW9UbqtA4pGGzoav7iSfkQs7ChcGJ9GMJzZeCBduWzyl7hzLB0q0N0zTesiJLGU6wni3ZaVuatQB_HgBsOyo6ERmr42tdyDXs5rgclW85j1NxcvSxrlctJFpUUZx4ZnUZcQw3xVeNuAZntINdnt1b8pWbw1M6A9F3ruEqpNimATWfopLWDCp5RdaCmva6McAh22wBjBiQ6GjWoBvKJhYg5KXXqpb2ojBkSDo3Y37buZZzm6LFlmQmSLRK4enMuymD-0Z17xRQYeMJdFBJ9XnGlyq2SXuURLa2KgxrVyW165SrYByEuVgPRoYH-RNkNWqFnZ7PGUyW2bYUWOU8L4BVITCq1ADBqMGslh2v-yHx2LKXN58R7sahUqpRae6LaMH39bY8tQabsOCwnwiiYl1cXjWwtVSRdaYKTS7a1GcDonqLAn2DjpdqATNSJYTvoXlFR_H1S0XyPR86S14Hen3pRvPFgb_0oG2sSlogeALEYH94wMYIqWgwAUEDw5lUL4XU6O94NVwckkUZ_6Q0oXWJkhs4uFofJRm23PKLcBa0k6i05y2VuFKNQljg7hzJe28bQ9FqyWKdP8ZTeFHqlpv_-uwSYwyCstIJxWspJRzKpuYViOeFVPYD6naF3mz2_eeFSF-lfHkfeJKckh62yK4pdhFs72OBiPkw4YgKSZKnXI8jAP-5Z4UMI1qDc0BPgghRcUEd9cPRQV1WrJ0gkMHfE3epGozgt3iDMot33pSHYf4rGpeK-soxI-rZ90znNyj2nOg7L74QMj-FrFCSad7V0x2aScg013jSQH7Z3ymWTCpK0Haam5N3S2A1qe8LlPqK4x5xyyoR6LXycszhBjxb2vad5aFzeRnxuxvBy5EC-AuNTUSlR-7QIqFQytHUg9nHIyp_MI5xsvPimrrgyYqi1VJANpuBk3me5JFSg_D4nMo9axhnOewU073de1SL_EU3KFDjkhXmQAjAjdCI5T0ayeC9TVMkOJ9IEAtbC7IIJhVR3BThe55y_9vrPaX3Lo5pHQkMGik5YfUqCPEV-Tk_m2jCa-tF-k_yaK_3fWBGchLFzGDotuYOML8YAq-AQ6gFQMP-9H7VAyPG6u4QHaXiTZX2Y56DBhC7vYvZfT_q6-cwf6zQrB-9U1oLiUDfuHxrwLdcb8yFcw",
		"https://*****.club/page/",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0",
	)
	preUrl := baseInfo.Url
	for i := 1; i <= 1101; i++ {
		baseInfo.Url = preUrl + strconv.Itoa(i)
		err := baseInfo.proxyGetPageOne(&pageInfoList, i)
		if err != nil {
			FairTime++
		}
		if FairTime > 30 {
			break
		}
		randNumber = time.Duration(rand.Intn(500) + 1000)
		fmt.Println("page :", i, "----pageInfoListLen :", len(pageInfoList), "----waitTime : ", randNumber, "ms")
		time.Sleep(randNumber * time.Millisecond)
	}

	err := rexcel.BufferToExcel(pageInfoList, buffer, "", "")
	if err != nil {
		log.Println(err)
	}
}

func (l *BaseInfo) proxyGetPageOne(pageInfoList *PageInfoList, page int) (_err error) {
	c := colly.NewCollector(
		colly.Async(false),
		colly.UserAgent(l.userAgent),
		//colly.AllowedDomains("https://*****.club"),
	)
	//c.SetRequestTimeout(10 * time.Second)
	c.SetCookies(l.Url, []*http.Cookie{l.InitCookie()})
	c.SetProxy("socks://127.0.0.1:7890")
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		_err = err
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML(".blog-entry-title", func(e *colly.HTMLElement) {
		find := e.DOM.Find("a")
		src, _ := find.Attr("href")
		*pageInfoList = append(*pageInfoList, &PageInfo{
			Page:  page,
			Title: find.Text(),
			Href:  "",
			Src:   src,
		})
		//log.Println(find.Text())
	})
	//c.Limit(&colly.LimitRule{
	//	DomainGlob:  "*",
	//	RandomDelay: 1 * time.Second, // 延时
	//})
	c.Visit(l.Url)
	c.Wait()
	return nil
}

func ProxyDownload(cookie, userAgent string) {
	var (
		pageInfoList = make(PageInfoList, 0)
	)
	// 自定义文件
	pageInfoList, err := Read(
		".xlsx",
		"",
	)
	if err != nil {
		return
	}

	baseInfo := Init(cookie,
		"",
		userAgent,
	)
	// 下载操作
	for _, info := range pageInfoList {
		if info.Page == 1 {
			baseInfo.ProxySave(info)
			randNumber := time.Duration(rand.Intn(1000))
			time.Sleep(randNumber * time.Millisecond)
		}
	}
}

func ProxyDownloadForMysql(cookie, userAgent string) {
	var (
		pageInfoList = make(PageInfoList, 0)
		structure    = structure.QueryPageInfoParam{}
	)
	connect, err := Connect()
	dao := NewPageInfoDao(
		connect)

	if err != nil {
		return
	}
	structure.Page = 1
	pageInfoList, _, err = dao.Search(context.Background(), structure)

	baseInfo := Init(cookie,
		"",
		userAgent,
	)
	// 下载操作
	for _, info := range pageInfoList {
		//if info.Page == 1 {
		baseInfo.ProxySave(info)
		randNumber := time.Duration(rand.Intn(300))
		time.Sleep(randNumber * time.Millisecond)
		//}
	}
}

func (l *BaseInfo) ProxySave(info *PageInfo) {
	var (
		path = "F:\\1111secret\\" + info.Title
	)
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent(l.userAgent),
	)
	c.SetCookies(info.Src, []*http.Cookie{l.InitCookie()})
	c.SetProxy("socks://127.0.0.1:7890")

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	// 判断路径是否存在
	_, err := os.ReadDir(path)
	if err != nil {
		// 不存在就创建
		err = os.MkdirAll(path, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	addOnHtmlForProxy(c, info)
	c.Visit(info.Src)
	c.Wait()
}

// 抽离判断
func addOnHtmlForProxy(c *colly.Collector, info *PageInfo) {
	c.OnHTML(".wp-block-image", func(e *colly.HTMLElement) {
		imgLabel := e.DOM.Find("img")
		//log.Println(imgLabel)
		srcAddr, _ := imgLabel.Attr("src")
		//log.Println(srcAddr)
		writeFlied(srcAddr, info.Title)
	})
}

// 写文件
func writeFlied(src, title string) {
	resp, err := http.Get(src)
	if err != nil {
		log.Println("Error fetching image:", err)
		return
	}
	defer resp.Body.Close()
	// 保存图片到本地
	fileName := strconv.Itoa(int(snowflack.NewCustomNode().GenerateID().Int64())) + ".jpg"
	log.Printf(fileName)
	outFile, err := os.Create(fmt.Sprintf("F:/1111secret/%s/%s", title, fileName))
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}
}
