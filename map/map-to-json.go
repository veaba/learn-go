package main

import "encoding/json"

func main() {

	arrayObject := map[string]string{
		"高一":  "高一",
		"高二":  "高二",
		"高三":  "高三",
		"初一":  "初一",
		"初二":  "初二",
		"初三":  "初三",
		"四年级": "小四",
		"小四":  "小四",
		"五年级": "小五",
		"小五":  "小五",
		"六年级": "小六",
		"小六":  "小六",
		"一年级": "小一",
		"二年级": "小一",
		"三年级": "小一",
	}
	arrayObjectData, _ := json.Marshal(arrayObject)
	println("===>", string(arrayObjectData))
	println("===>", arrayObjectData)
}
