package alivod


const(
    //接口地址
    ALIVOD_API = "http://vod.cn-shanghai.aliyuncs.com"

    //上传
    ALIVOD_API_CREATE_UPLOAD_VIDEO = "CreateUploadVideo"
    ALIVOD_API_REFRESH_UPLOAD_VIDEO = "RefreshUploadVideo"
    ALIVOD_API_CREATE_UPLOAD_IMAGE = "CreateUploadImage"
    //播放
    ALIVOD_API_GET_PLAY_INFO = "GetPlayInfo"
    ALIVOD_API_GET_VIDEO_PLAY_AUTH = "GetVideoPlayAuth"
    //视频管理
    ALIVOD_API_GET_VIDEO_INFO = "GetVideoInfo"
    ALIVOD_API_UPDATE_VIDEO_INFO = "UpdateVideoInfo"
    ALIVOD_API_DELETE_VIDEO = "DeleteVideo"
    ALIVOD_API_GET_VIDEO_LIST = "GetVideoList"
    //分类
    ALIVOD_API_ADD_CATEGORY = "AddCategory"
    ALIVOD_API_UPDATE_CATEGORY = "UpdateCategory"
    ALIVOD_API_DELETE_CATEGORY = "DeleteCategory"
    ALIVOD_API_GET_CATEGORIES = "GetCategories"

)


type AlivodConfig struct {
    //↓↓↓↓↓↓↓↓↓↓请在这里配置您的基本信息↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
    //阿里视频点播Access key id
    AccessKeyId string
    //阿里视频点播Access key secret
    AccessKeySecret string
    //sdk版本
    AlivodVersion  string
    //阿里视频点播接口地址
    AlivodAPI string
    //返回值类型,支持JSON与XML，默认为XML
    AlivodFormat string
    //签名方式，目前支持HMAC-SHA1。
    AlivodSignatureMethod string
    //签名算法版本，目前版本是1.0。
    AlivodSignatureVersion string

}

var _alivodConfig = AlivodConfig {
    //fot test
    AccessKeyId: "testId",
    AccessKeySecret: "testKeySecret",

    AlivodAPI: ALIVOD_API,
    AlivodVersion: "2017-03-21",
    AlivodFormat: "JSON",
    AlivodSignatureMethod: "HMAC-SHA1",
    AlivodSignatureVersion: "1.0",
}

/**
 * 初始化用户access key id, access key secret
 */
func InitConfig(accessKeyId, accessKeySecret string) {
    _alivodConfig.AccessKeySecret = accessKeySecret
    _alivodConfig.AccessKeyId = accessKeyId
}

func GetAlivodConfig() *AlivodConfig {
    return &_alivodConfig
}
