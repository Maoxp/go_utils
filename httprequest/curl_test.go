package httprequest

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpPost(t *testing.T) {
	client := http.Client{}
	head := http.Header{}
	head.Set("Authorization", "AZXddWTGSFC4FSAJXNGSCG47ASFDLDHEBCH")
	head.Set("Content-type", "application/json")
	//head.Set("Content-type", "application/x-www-form-urlencoded")
	target := "https://xxxxxx.com/v2/it/category"

	// strings.NewReader() 用法：
	/*var r http.Request
	_ = r.ParseForm()
	r.Form.Add("page", "2")
	bodyStr := strings.TrimSpace(r.Form.Encode())
	bodyIo, err := HttpPostForm(&client, target, head, bodyStr)*/

	//bodyIo, err := HttpPost(&client, target, head, []byte("page=1&pageSize=10&keyword=红叶石楠"))
	bodyIo, err := HttpPost(&client, target, head, []byte("{\"page\":1, \"pageSize\":10, \"keyword\":\"红叶石楠\"}"))
	if err != nil {
		t.Fatal("error:", err)
	}
	bodyByte, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		t.Fatal("ReadAll error:", err)
	}
	t.Fatal(string(bodyByte))
}

func TestHttpGet(t *testing.T) {
	client := http.Client{}
	head := http.Header{}
	head.Set("Authorization", "AZXddWTGSFC4FSAJXNGSCG47ASFDLDHEBCH")
	head.Set("Content-type", "application/json")
	target := "https://xxxxxx.com/v2/it/category"
	bodyIo, err := HttpGet(&client, target, head)
	if err != nil {
		t.Fatal("error:", err)
	}
	defer bodyIo.Close()
	bodyByte, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		t.Fatal("ReadAll error:", err)
	}
	t.Fatal(string(bodyByte))
}
