/*
 * Package : codeVarGo
 * Author : zhongbr
 * CreateTime :
 * @ All rights reserved .
 */
package variable

import (
	"regexp"
	"strings"
)

// 解析输入，返回操作类型和结果
func ParseAction(input string) (string, string) {
	// 检查输入是否包含汉字
	chineseChecker := regexp.MustCompile("[\u4e00-\u9fa5]")
	specialChecker := regexp.MustCompile("[^\u4e00-\u9fa5A-Za-z0-9_$]")
	if chineseChecker.MatchString(input) {
		if specialChecker.MatchString(input) {
			return "注释汉译英", input
		} else {
			return "变量汉译英", input
		}
	} else {
		if specialChecker.MatchString(input) {
			return "注释英译汉", input
		} else {
			return "变量英译汉", VariNameTurning(input)
		}
	}
}

func VariNameTurning(input string) string {
	input = ParseXh(input)
	if regexp.MustCompile("[a-z]").MatchString(input) {
		input = ParseTf(input)
	} else {
		input = strings.ToLower(input)
	}
	return input
}

func ParseTf(input string) string {
	const uppers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for char := range uppers {
		upper := string(uppers[char])
		input = strings.ReplaceAll(input, upper, " "+strings.ToLower(upper))
	}
	if strings.HasPrefix(input, " ") {
		input = input[1:]
	}
	return input
}

func ParseXh(input string) string {
	return strings.ReplaceAll(input, "_", " ")
}

// 下划线法
func ToXh(input string) string {
	input = strings.ToLower(input)
	input = RemoveUnecessaryWords(input)
	input = regexp.MustCompile("[^a-z0-9_]").ReplaceAllString(input, " ")
	input = strings.ReplaceAll(input, " ", "_")
	input = regexp.MustCompile("_+").ReplaceAllString(input, "_")
	input = regexp.MustCompile("^_").ReplaceAllString(input, "")
	return input
}

// 驼峰法
func ToTf(input string) string {
	input = strings.ToLower(input)
	input = RemoveUnecessaryWords(input)
	input = regexp.MustCompile("[^a-z0-9 ]").ReplaceAllString(input, " ")
	words := strings.Split(input, " ")
	result := words[0]
	for index := range words {
		if index != 0 {
			if len(words[index]) > 1 {
				result += strings.ToUpper(string(words[index][0])) + words[index][1:]
			} else {
				result += strings.ToUpper(words[index])
			}
		}
	}
	return result
}

// 常量法
func ToConst(input string) string {
	input = ToXh(input)
	return strings.ToUpper(input)
}

// 去掉虚词
func RemoveUnecessaryWords(input string) string {
	input = regexp.MustCompile("^the | the ").ReplaceAllString(input, " ")
	input = regexp.MustCompile("^a | a |^an | an ").ReplaceAllString(input, " ")
	input = regexp.MustCompile("^ ").ReplaceAllString(input, "")
	return input
}
