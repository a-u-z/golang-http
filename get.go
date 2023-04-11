package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r *Req) get(path string) (respString string, err error) {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, r.baseUrl+path, nil)
	if err != nil {
		fp, n, l := getFilePathFuncNameAndLine()
		err = fmt.Errorf(fp, n, l, err)
		return
	}

	// 設置 HTTP 標頭
	for _, v := range r.header {
		req.Header.Add(v.key, fmt.Sprintf("%v", v.value))
	}

	resp, err := client.Do(req)
	if err != nil {
		fp, n, l := getFilePathFuncNameAndLine()
		err = fmt.Errorf(fp, n, l, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fp, n, l := getFilePathFuncNameAndLine()
		err = fmt.Errorf(fp, n, l, err)
		return
	}

	return string(body), nil
}
