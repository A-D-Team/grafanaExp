package internal

import (
	"crypto/tls"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"penTools/grafanaExp/pkg/http"
)

func DoReq(_url string) (re string) {
	//Logger.Debug(_url)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err := http.NewRequest("GET", _url, nil)
	c := &http.Client{Transport: tr}
	req.PathAsIs.Flag = true
	req.PathAsIs.RawUrl = _url
	resp, err := c.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
