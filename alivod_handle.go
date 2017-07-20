package alivod

import (

)
import (
    "bytes"
    "net/url"
)

func getCommonKvpairs(alivodConfig *AlivodConfig) Kvpairs {
    return Kvpairs{
        //公共参数部分
        Kvpair{`AccessKeyId`, alivodConfig.AccessKeyId},
        Kvpair{`Format`, alivodConfig.AlivodFormat},
        Kvpair{`Version`, alivodConfig.AlivodVersion},
        Kvpair{`SignatureMethod`, alivodConfig.AlivodSignatureMethod},
        Kvpair{`SignatureVersion`, alivodConfig.AlivodSignatureVersion},
        Kvpair{`SignatureNonce`, getNonceStr()},
        Kvpair{`Timestamp`, getDateTime()},
        //Kvpair{`Signature`, },        //待计算部分
    }
}

/**
 * 获取视频上传地址和凭证，并创建视频信息。
 * CreateUploadVideo处理
 */
func GetAlivodCreateUploadVideoUrl(acuv CreateUploadVideoRequest) string {
    //获取公共参数
    param := getCommonKvpairs(GetAlivodConfig())
    //设置参数
    param = append(param, Kvpair{`Action`, acuv.Action})
    param = append(param, Kvpair{`Title`, acuv.Title})
    param = append(param, Kvpair{`FileName`, acuv.FileName})
    param = append(param, Kvpair{`FileSize`, acuv.FileSize})
    param = append(param, Kvpair{`Description`, acuv.Description})
    param = append(param, Kvpair{`CoverURL`, acuv.CoverURL})
    param = append(param, Kvpair{`CateId`, acuv.CateId})
    param = append(param, Kvpair{`Tags`, acuv.Tags})
    uri := hmacSha1Sign(&param, GetAlivodConfig())
    return getRequestUrl(uri, GetAlivodConfig())
}


/**
 * 获取图片上传地址和凭证
 * CreateUploadImage处理
 */
func GetAlivodCreateUploadImageUrl(acui CreateUploadImageRequest) string {
    //获取公共参数
    param := getCommonKvpairs(GetAlivodConfig())
    //设置参数
    param = append(param, Kvpair{`Action`, acui.Action})
    param = append(param, Kvpair{`ImageType`, acui.ImageType})
    param = append(param, Kvpair{`ImageExt`, acui.ImageExt})
    uri := hmacSha1Sign(&param, GetAlivodConfig())
    return getRequestUrl(uri, GetAlivodConfig())
}

/**
 * 获取视频上传地址和凭证，并创建视频信息。
 * GetVideoPlayAuth处理
 */
func GetAlivodGetVideoPlayAuthUrl(acvpa GetVideoPlayAuthRequest) string {
    //获取公共参数
    param := getCommonKvpairs(GetAlivodConfig())
    //设置参数
    param = append(param, Kvpair{`Action`, acvpa.Action})
    param = append(param, Kvpair{`VideoId`, acvpa.VideoId})
    uri := hmacSha1Sign(&param, GetAlivodConfig())
    return getRequestUrl(uri, GetAlivodConfig())
}



/**
 * 获取视频信息
 * GetVideoInfo
 */
func GetAlivodVideoInfoUrl(agvi GetVideoInfoRequest) string {
    //获取公共参数
    param := getCommonKvpairs(GetAlivodConfig())
    //设置参数
    param = append(param, Kvpair{`Action`, agvi.Action})
    param = append(param, Kvpair{`VideoId`, agvi.VideoId})
    uri := hmacSha1Sign(&param, GetAlivodConfig())
    return getRequestUrl(uri, GetAlivodConfig())
}



func hmacSha1Sign(para *Kvpairs, alivodConfig *AlivodConfig) string {
    buildRequestPara(para, alivodConfig)
    return createLinkstringUrlencode(para)
}




func getRequestUrl(params string, alivodConfig *AlivodConfig) string {
    var buf bytes.Buffer
    defer buf.Reset()
    buf.WriteString(alivodConfig.AlivodAPI)
    buf.WriteString("?")
    buf.WriteString(params)
    uri, _ := url.PathUnescape(buf.String())
    return uri
}