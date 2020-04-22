/*
 * Package : main
 * Author : zhongbr
 * CreateTime :
 * @ All rights reserved .
 */
package test

import (
	"alfred-var/variable"
	"encoding/json"
	"fmt"
	"testing"
)

const src = `{
  "type": "ZH_CN2EN",
  "errorCode": 0,
  "elapsedTime": 17,
  "translateResult": [
    [
      {
        "src": "{处理后等待翻译的变量文本}",
        "tgt": "{} variable text waiting for translation in"
      }
    ]
  ]
}`

func TestJsonFormat(t *testing.T) {
	var translation map[string]interface{}
	err := json.Unmarshal([]byte(src), &translation)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		result := translation["translateResult"].([]interface{})[0].([]interface{})[0].(map[string]interface{})["tgt"].(string)
		fmt.Println(result)
	}
}

func TestReadAbb(t *testing.T) {
	fmt.Println(variable.Abb("electronic data interchange"))
}
