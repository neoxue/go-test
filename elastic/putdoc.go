package main

import (
	//"net/http"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

//import "github.com/stretchr/stew/objects"
import (
	//"github.com/NodePrime/jsonpath/cli/jsonpath"
	//"github.com/NodePrime/jsonpath"
	//"github.com/stretchr/stew/objects"
	//"github.com/NodePrime/jsonpath"
	//"gopl.io/ch7/eval"
	"os"
	"strings"
)

func main() {
	// 暂时先写死，之后改用合适的 jsonpath

	//var docArr = [20]string {"fymvece1601193"}
	var docArr = [20]string {"fymvece1601193", "fxmpnqf9628055"}
	for i := 0; i < 2; i++ {
		resp, _ := http.Get("http://ds.pub.sina.com.cn:8080/api/doc/get?docID=" + docArr[i])
		body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
		//var m 	= new(struct{result struct{data struct{items []map[string] interface{}}}})
		//var m map[string] map[string] map[string] []map[string] interface{}
		var m map[string] map[string] map[string] interface{}
		var err = json.Unmarshal([]byte(body), &m)
		if err == nil {
			//fmt.Println(m["result"]["data"]["items"])
			m2 := m["result"]["data"]["items"].([]interface{});
			m3 := m2[0]
			m4 := m3.(map[string]interface{})
			//fmt.Println(m4)
			id := m4["_id"]
			delete(m4, "_id")
			//fmt.Println(id)
			docByteArr, _ := json.Marshal(m4);
			// 写入文件 利用 bulk 方式添加

			docAddInfo := `{ "index" : { "_index" : "acomos_v2", "_type" : "comos", "_id" : "`;
			docAddInfo += id.(string)
			docAddInfo += `" }}`
			writeFile([] byte(docAddInfo))
			writeFile([] byte("\n"))
			writeFile(docByteArr)
			writeFile([] byte("\n"))
			writeFile([] byte("\n"))
		}
	}

	push();
}

func push() {
	client := &http.Client{}
	postStr, _ := ioutil.ReadFile("/tmp/comos.txt")
	req, err := http.NewRequest("POST", "http://localhost:9200/_bulk?pretty", strings.NewReader(string(postStr)))
	if err != nil {
		fmt.Println("err" , err);
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err);
	}
	fmt.Println(string(body))
}

func writeFile(docByteArr []byte) {
	fd,err:=os.OpenFile("/tmp/comos.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		fmt.Println(err)
	}
	fd.Write(docByteArr)
	fd.Close()
}

