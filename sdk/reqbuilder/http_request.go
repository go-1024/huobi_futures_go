package reqbuilder

import (
	"github.com/gostudys/huobi_futures_go/sdk/log"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) (string, error) {
	logger := log.GetInstance()
	logger.Start()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("GET", url)

	return string(result), err
}

func HttpPost(url string, body string) (string, error) {
	logger := log.GetInstance()
	logger.Start()

	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("POST", url)

	return string(result), err
}
