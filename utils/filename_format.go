package utils

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
	"time"
)

const FormatRE = "{{.+?}}"

type NameFormat string

func (n NameFormat) GetName() string {
	tmp := strings.TrimLeft(string(n), "{{")
	tmp = strings.TrimRight(tmp, "}}")
	tmp = strings.Trim(tmp, " ")
	if strings.Contains(tmp, "|") {
		t := strings.Split(tmp, "|")
		return n.MethodWithParams(t[0], strings.Split(t[1], ",")...)
	}
	return n.Method(tmp)
}

func (n NameFormat) Method(t string) string {
	switch t {
	case "md5":
		return MD5(time.Now().String())
	default:
		return t
	}
}

func (n NameFormat) MethodWithParams(t string, p ...string) string {
	switch t {
	case "date":
		return time.Now().Format(p[0])
	default:
		return t
	}
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Replace(str string) string {
	re := regexp.MustCompile(FormatRE)
	out := re.FindAllString(str, -1)
	//fmt.Println(out)
	if out == nil {
		f := strings.Split(str, "/")
		return f[len(f)-1]
	}
	//如果没有匹配到，按原来的格式返回
	if len(out) == 0 {
		return str
	}

	s := re.Split(str, -1)
	arr := make([]string, 0)
	for i, o := range s {
		arr = append(arr, o)
		if i < len(s)-1 {
			name := NameFormat(out[i])
			arr = append(arr, name.GetName())
		}
	}
	//fmt.Println(arr)

	return strings.Join(arr, "")
}
