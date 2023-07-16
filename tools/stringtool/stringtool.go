package stringtool

import (
	"github.com/google/uuid"
	"strings"
)

var numberDict = map[string]int64{"一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9, "零": 0, "壹": 1, "贰": 2, "叁": 3, "肆": 4, "伍": 5, "陆": 6, "柒": 7, "捌": 8, "玖": 9, "两": 2}

var unitDict = map[string]int64{"千": 1000, "百": 100, "十": 10}

var largeUnitDict = map[string]int64{"万": 10000, "亿": 100000000}

// ChineseNumStringToArabicNum 将中文数字转换为阿拉伯数字, 仅支持到亿,只支持整数
func ChineseNumStringToArabicNum(s string) (res int64) {
	var part, tmp int64
	for _, ele := range s {
		if val, ok := numberDict[string(ele)]; ok {
			tmp = val
			continue
		}
		if val, ok := unitDict[string(ele)]; ok {
			if val == 10 && tmp == 0 {
				tmp = 1
			}
			part += tmp * val
			tmp = 0
			continue
		}
		if val, ok := largeUnitDict[string(ele)]; ok {
			part += tmp
			tmp = 0
			part = part * val
			if string(ele) == "亿" {
				res += part
				part = 0
			}
		}
	}
	res = res + part + tmp
	return
}

func GenerateUUID() string {
	uuid := uuid.New()
	return strings.ReplaceAll(uuid.String(), "-", "")
}
