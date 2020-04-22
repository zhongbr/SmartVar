/*
 * Package : alfred_var
 * Author : zhongbr
 * CreateTime :
 * @ All rights reserved .
 */
package main

import (
	"alfred-var/variable"
	yd_translate "alfred-var/yd-translate"
	"flag"
	aw "github.com/deanishe/awgo"
	"strings"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func capitalizeFirstLetter(tf string) string {
	var capitalize string
	if len(tf) > 1 {
		capitalize = strings.ToUpper(string(tf[0])) + tf[1:]
	} else {
		capitalize = strings.ToUpper(tf)
	}
	return capitalize
}

// 缩写
func AbbTransform(text string, wf *aw.Workflow) {
	// 尝试缩写
	abb := variable.Abb(text)
	if abb != text {
		wf.NewItem("缩写-驼峰法:" + variable.ToTf(abb)).Subtitle("缩写变形").Valid(true).Arg(variable.ToTf(abb))
		wf.NewItem("缩写-下划线:" + variable.ToXh(abb)).Subtitle("缩写变形").Valid(true).Arg(variable.ToXh(abb))
		wf.NewItem("缩写-常量法:" + variable.ToConst(abb)).Subtitle("缩写变形").Valid(true).Arg(variable.ToConst(abb))
		wf.NewItem("缩写-大写驼峰法:" + capitalizeFirstLetter(variable.ToTf(abb))).Subtitle("缩写变形").Valid(true).Arg(capitalizeFirstLetter(variable.ToTf(abb)))
	}
	// 尝试还原缩写
	prototype := variable.Prototype(text)
	if prototype != text {
		wf.NewItem("还原-驼峰法:" + variable.ToTf(prototype)).Subtitle("还原变形").Valid(true).Arg(variable.ToTf(prototype))
		wf.NewItem("还原-下划线:" + variable.ToXh(prototype)).Subtitle("还原变形").Valid(true).Arg(variable.ToXh(prototype))
		wf.NewItem("还原-常量法:" + variable.ToConst(prototype)).Subtitle("还原变形").Valid(true).Arg(variable.ToConst(prototype))
		wf.NewItem("还原-大写驼峰法:" + capitalizeFirstLetter(variable.ToTf(prototype))).Subtitle("还原变形").Valid(true).Arg(capitalizeFirstLetter(variable.ToTf(prototype)))
	}
}

// 解析参数
func argsParse(args []string, workflow *aw.Workflow) {
	input := ""
	for i := range args {
		input += args[i]
		if i != len(args)-1 {
			input += " "
		}
	}
	action, text := variable.ParseAction(input)
	// 尝试还原缩写
	prototype := variable.Prototype(text)
	translation, err := yd_translate.Translate(prototype)
	if err != nil {
		wf.NewItem("翻译错误了哦！" + err.Error())
	} else {
		var result string
		if action == "变量英译汉" {
			result = translation
			wf.NewItem("变量名含义:" + result).Subtitle("英语短语：" + prototype).Valid(true).Arg(result)
			// 格式转换
			wf.NewItem("驼峰法:" + variable.ToTf(text)).Subtitle("变量格式转换").Arg(variable.ToTf(text)).Valid(true)
			wf.NewItem("下划线:" + variable.ToXh(text)).Subtitle("变量格式转换").Arg(variable.ToXh(text)).Valid(true)
			wf.NewItem("常量法:" + variable.ToConst(text)).Subtitle("变量格式转换").Arg(variable.ToConst(text)).Valid(true)
			var capitalizeFirstLetter string
			tf := variable.ToTf(text)
			if len(tf) > 1 {
				capitalizeFirstLetter = strings.ToUpper(string(tf[0])) + tf[1:]
			} else {
				capitalizeFirstLetter = strings.ToUpper(tf)
			}
			wf.NewItem("首字母大写:" + capitalizeFirstLetter).Subtitle("变量格式转换").Arg(capitalizeFirstLetter).Valid(true)
			// 尝试缩写和还原
			AbbTransform(text, wf)
		} else if action == "变量汉译英" {
			var value string
			value = variable.ToTf(translation)
			wf.NewItem("驼峰法:" + value).Subtitle(value).Valid(true).Arg(value)
			value = variable.ToXh(translation)
			wf.NewItem("下划线:" + value).Subtitle(value).Valid(true).Arg(value)
			value = variable.ToConst(translation)
			wf.NewItem("常量法:" + value).Subtitle(value).Valid(true).Arg(value)
			var capitalizeFirstLetter string
			tf := variable.ToTf(translation)
			if len(tf) > 1 {
				capitalizeFirstLetter = strings.ToUpper(string(tf[0])) + tf[1:]
			} else {
				capitalizeFirstLetter = strings.ToUpper(tf)
			}
			wf.NewItem("大写驼峰法:" + capitalizeFirstLetter).Subtitle("变量格式转换").Arg(capitalizeFirstLetter).Valid(true)
			// 尝试缩写和还原
			AbbTransform(translation, wf)
		} else if action == "注释英译汉" {
			wf.NewItem(action).Subtitle(translation).Valid(true).Arg(translation)
		} else if action == "注释汉译英" {
			wf.NewItem(action).Subtitle(translation).Valid(true).Arg(translation)
		}
	}
}

func run() {
	// 解析参数
	flag.Parse()
	argsParse(flag.Args(), wf)
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
