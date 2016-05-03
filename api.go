package b9

import (
	"errors"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Info struct {
	Title  string `json:"title"`
	Src    string `json:"src"`
	ImgURL string `json:"img_url"`
	Time   string `json:"time"`
}

func GetNormal(n int) ([]Info, error) {
	url := "http://up.b9dm.com/index.php/video/show/cid/5/order/2/page/" + strconv.Itoa(n)
	return getInfo(url)
}

func GetHD(n int) ([]Info, error) {
	url := "http://up.b9dm.com/index.php/video/show/cid/56/order/2/page/" + strconv.Itoa(n)
	return getInfo(url)
}

func getInfo(url string) ([]Info, error) {
	var ret []Info
	doc, err := goquery.NewDocument(url)
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
		buf.Title = src.Text()
		buf.Time = dd.Eq(1).Text()
		ret = append(ret, buf)
	})
	if len(ret) == 0 {
		return nil, errors.New("no data")
	}
	return ret, nil
}
