/*
 * Package : variable
 * Author : zhongbr
 * CreateTime :
 * @ All rights reserved .
 */
package variable

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func ReadAbbreviations() map[string]string {
	bytes, err := ioutil.ReadFile("./abbreviations.json")
	if err != nil {
		return nil
	}
	var result map[string]string
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil
	}
	return result
}

func Prototype(input string) string {
	words := strings.Split(input, " ")
	abb := ReadAbbreviations()
	for index := range words {
		prototype, exist := abb[words[index]]
		if exist {
			input = strings.ReplaceAll(input, words[index], prototype)
		}
	}
	return input
}

func Abb(input string) string {
	abbs := ReadAbbreviations()
	for abb := range abbs {
		prototype := abbs[abb]
		input = strings.ReplaceAll(input, prototype, abb)
	}
	return input
}
