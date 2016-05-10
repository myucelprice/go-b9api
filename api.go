package b9

import (
	"errors"
	"net/http"
	"strconv"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

type API struct {
	UA string
}

type Info struct {
	Title      string `json:"title"`
	ShortTitle string `json:"short_title"`
	Src        string `json:"src"`
	ImgURL     string `json:"img_url"`
	Time       string `json:"time"`
}

func NewAPI() *API {
	return &API{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:46.0) Gecko/20100101 Firefox/46.0",
	}
}

func (api *API) SetUA(ua string) {
	api.UA = ua
}

func (api *API) GetNormal(n int) ([]Info, error) {
	url := "http://up.b9dm.com/index.php/video/show/cid/5/order/2/page/" + strconv.Itoa(n)
	return getInfo(url, api.UA)
}

func (api *API) GetHD(n int) ([]Info, error) {
	url := "http://up.b9dm.com/index.php/video/show/cid/56/order/2/page/" + strconv.Itoa(n)
	return getInfo(url, api.UA)
}

func getInfo(url string, ua string) ([]Info, error) {
	var ret []Info
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	root, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	doc := goquery.NewDocumentFromNode(root)

	if err != nil {
		return nil, err
	}
	doc.Find("dl.t_box").Each(func(i int, s *goquery.Selection) {
		var buf Info
		dt := s.Find("dt")
		dd := s.Find("dd")
		a := dt.Find("a")
		img := a.Eq(0).Find("img")
		src := a.Eq(1)
		buf.ImgURL, _ = img.Attr("src")
		buf.Src, _ = src.Attr("href")
		buf.Title, _ = src.Attr("title")
		buf.ShortTitle = src.Text()
		buf.Time = dd.Eq(1).Text()
		ret = append(ret, buf)
	})
	if len(ret) == 0 {
		return nil, errors.New("no data")
	}
	return ret, nil
}
