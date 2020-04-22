/*
 * Package : yd_translate
 * Author : zhongbr
 * CreateTime :
 * @ All rights reserved .
 */
package yd_translate

import (
	"encoding/json"
	"fmt"
	"github.com/Unknwon/com"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Translate(input string) (string, error) {
	var url, result string
	client := http.Client{}
	chineseChecker := regexp.MustCompile("[\u4e00-\u9fa5]")
	if chineseChecker.MatchString(input) {
		url = fmt.Sprintf("http://fanyi.youdao.com/translate?&doctype=json&type=ZH_CN2EN&i=%s", com.UrlEncode(input))
	} else {
		url = fmt.Sprintf("http://fanyi.youdao.com/translate?&doctype=json&type=EN2ZH_CN&i=%s", com.UrlEncode(input))
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	var translation map[string]interface{}
	err = json.Unmarshal(bytes, &translation)
	if err != nil {
		return "", err
	}
	result = translation["translateResult"].([]interface{})[0].([]interface{})[0].(map[string]interface{})["tgt"].(string)
	return result, nil
}
