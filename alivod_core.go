package alivod

import (
    "time"
    "math/rand"
    "io/ioutil"
    "net/http"
    "fmt"
    "crypto/tls"
)


/**
 * 获取alivod 时间
 *请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ
例如，2017-3-29T12:00:00Z(为北京时间2017年3月29日的20点0分0秒。
 */

func getDateTime() string {
    fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
    return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func getRandomString(length int) string {
    if length <= 0 {
        return ""
    }
    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    rawString := []byte(str)
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    result := []byte{}
    for j := 0; j < length; j++ {
        result = append(result, rawString[r.Intn(len(rawString))])
    }
    return string(result)
}

func getNonceStr() string {
    return getRandomString(12)
}



func GetHttpResponse(url string) (string, error) {
    fmt.Println("http.Client get url : ", url)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("http.Client get err : ", err)
        return "", err
    }
    bodyByte, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("http.Client get bodyByte : ", string(bodyByte))
    return string(bodyByte), nil
}

/**
 * 实现多种字符编码方式
 * @param $input 需要编码的字符串
 * @param $_output_charset 输出的编码格式
 * @param $_input_charset 输入的编码格式
 * return 编码后的字符串
 */
func charsetEncode() {}

/**
 * 实现多种字符解码方式
 * @param $input 需要解码的字符串
 * @param $_output_charset 输出的解码格式
 * @param $_input_charset 输入的解码格式
 * return 解码后的字符串
 */
func charsetDecode() {}

/**
 * 生成签名结果
 * @param $para_sort 已排序要签名的数组
 * return 签名结果字符串
 */
func buildRequestMysign(kv *Kvpairs, alivodConfig *AlivodConfig) string {
	mysign := ""
	switch alivodConfig.AlivodSignatureMethod {
	case "HMAC-SHA1":
		prestr := createLinkstringUrlencode(kv)
		mysign = sha1Sign(prestr, alivodConfig.AccessKeySecret)
		break
	//case "MD5":
	//	return md5Sign(createLinkStringNoUrl(kv), alipayConfig.Key)
	//	break
	default:
	}
	return mysign
}

/**
 * 视频点播签名算法
 * 生成要请求给阿里云视频点播的参数数组
 * @param $para_temp 请求前的参数数组
 * @return 要请求的参数数组
 * @link https://help.aliyun.com/document_detail/44434.html?spm=5176.doc55407.6.598.v1RbFe
 */
func buildRequestPara(kv *Kvpairs, alivodConfig *AlivodConfig) {
	//除去待签名参数数组中的空值和签名参数
	paraFilter(kv)
	//对待签名参数数组排序
	argSort(kv)
	//生成签名结果
	mysign := buildRequestMysign(kv, alivodConfig)
	//签名结果与签名方式加入请求提交参数组中
	*kv = append(*kv, Kvpair{`Signature`, mysign})
}


