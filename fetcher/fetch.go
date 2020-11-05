package fetcher

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100*time.Millisecond)

func Fetch(url string)([]byte,error)  {
	<-rateLimiter
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent",getAgent())
	req.Header.Set("Connection", "keep-alive")
	resp, err := (&http.Client{

	}).Do(req)
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		if resp.StatusCode == http.StatusAccepted{
			resp, err = (&http.Client{}).Do(req)
		}else {
			return nil,fmt.Errorf("Error:status error:%d ",resp.StatusCode)
		}
	}

	respReader := bufio.NewReader(resp.Body)
	e := determineEncoding(respReader)
	utf8Reader:=transform.NewReader(respReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader)encoding.Encoding  {

	bytes,err :=r.Peek(1024)
	if err != nil{
		logs.Info("Fetcher error: %v",err)
		return unicode.UTF8
	}
	e,_, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func getAgent() string {
	agent  := [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(agent)
	return agent[r.Intn(len)]
}
