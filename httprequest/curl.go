package httprequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

var UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1541.0 Safari/537.36"

// strings.NewReader()、bytes.NewReader()、bytes.NewBuffer() 均使用了reader接口
func httpCall(client *http.Client, method, url string, header http.Header, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	for k, vs := range header {
		req.Header[k] = vs
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		//bodyByte, err := ioutil.ReadAll(resp.Body)
		return resp.Body, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 { // 403 can be rate limit error.  || resp.StatusCode == 403 {
		err = fmt.Errorf("resource not found: %s", url)
	} else {
		err = fmt.Errorf("%s %s -> %d", method, url, resp.StatusCode)
	}
	return nil, err
}

/*
params := url.Values{}
params.Set("code", code)
//如果参数中有中文参数,这个方法会进行URLEncode
Url.RawQuery = params.Encode()
urlPath := Url.String()
*/
func HttpGet(client *http.Client, url string, header http.Header) (io.ReadCloser, error) {
	return httpCall(client, "GET", url, header, nil)
}

func HttpPost(client *http.Client, url string, header http.Header, body []byte) (io.ReadCloser, error) {
	return httpCall(client, "POST", url, header, bytes.NewBuffer(body))
}

// Content-Type为application/x-www-form-urlencoded
//func HttpPostForm(client *http.Client, url string, header http.Header, body string) (io.ReadCloser, error) {
//	return httpCall(client, "POST", url, header, strings.NewReader(body))
//}



// HttpGetToFile gets the specified resource and writes to file.
// ErrNotFound is returned if the server responds with status 404.
func HttpGetToFile(client *http.Client, url string, header http.Header, fileName string) error {
	rc, err := HttpGet(client, url, header)
	if err != nil {
		return err
	}
	defer rc.Close()

	os.MkdirAll(path.Dir(fileName), os.ModePerm)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, rc)
	return err
}


func HttpGetBytes(client *http.Client, url string, header http.Header) ([]byte, error) {
	rc, err := HttpGet(client, url, header)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}

// HttpGetJSON gets the specified resource and mapping to struct.
// ErrNotFound is returned if the server responds with status 404.
func HttpGetJSON(client *http.Client, url string, header http.Header, v interface{}) error {
	rc, err := HttpGet(client, url, header)
	if err != nil {
		return err
	}
	defer rc.Close()
	err = json.NewDecoder(rc).Decode(v)
	if _, ok := err.(*json.SyntaxError); ok {
		return fmt.Errorf("JSON syntax error at %s", url)
	}
	return nil
}