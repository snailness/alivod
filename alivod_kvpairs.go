package alivod

import (
	"net/url"
	"sort"
	"strings"
)

type Kvpair struct {
	K, V string
}

type Kvpairs []Kvpair

func (t Kvpairs) Less(i, j int) bool {
	return t[i].K < t[j].K
}

func (t Kvpairs) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Kvpairs) Len() int {
	return len(t)
}

func (t Kvpairs) Sort() {
	sort.Sort(t)
}

/**
 * 除去数组中的空值和签名参数
 * @param $para 签名参数组
 * return 去掉空值与签名参数后的新签名参数组
 */
func paraFilter(para *Kvpairs) {
	new := Kvpairs{}
	for _, kv := range *para {
		if kv.V != "" {
			new = append(new, kv)
		}
	}
	*para = new
}

/**
 * 对数组排序
 * @param $para 排序前的数组
 * return 排序后的数组
 */
func argSort(para *Kvpairs) {
	para.Sort()
}

/**
 * 把数组所有元素，按照“参数=参数值”的模式用“&”字符拼接成字符串，并对字符串做urlencode编码
 * @param $para 需要拼接的数组
 * return 拼接完成以后的字符串
 */

func createLinkstringUrlencode(para *Kvpairs) string {
	var strs []string
	for _, kv := range *para {
        strs = append(strs, kv.K+"=" + urlEscape(kv.V))
	}
	return urlEscape(strings.Join(strs, "&"))
}

func createLinkstringForPost(para *Kvpairs) string {
    var strs []string
    for _, kv := range *para {
        strs = append(strs, kv.K+"=" + urlEscape(kv.V))
    }
    return urlEscape(strings.Join(strs, "&"))
}


/**
 * 一般支持URL编码的库（比如Java中的java.net.URLEncoder）都是按照”application/x-www-form-urlencoded”的MIME类型的规则进行编码的。实现时可以直接使用这类方式进行编码，把编码后的字符串中加号（+）替换成%20、星号（*）替换成%2A、%7E替换回波浪号(~)，即可得到上述规则描述的编码字符串。
 */
func urlEscape(s string) string {
    if len(s) > 0 {
        s = url.QueryEscape(s)
        //特殊处理
        s = strings.Replace(s, "+", "%20", -1)
        s = strings.Replace(s, "*", "%2A", -1)
        s = strings.Replace(s, "%7E", "~", -1)
    }
    return s
}