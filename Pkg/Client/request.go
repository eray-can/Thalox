package Client

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func CreateRequest(url string, method string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	headerSet(req, headers)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Response(response *http.Response) ([]byte, error) {

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func headerSet(req *http.Request, headers map[string]string) {

	req.Header.Set("Connection", "\"keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("Connection", "keep-alive")

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
