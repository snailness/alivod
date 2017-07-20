package alivod


/**
 * 公共参数
 */
type AlivodRequest struct {
    Format string
    Version    string
    AccessKeyId   string
    Signature       string
    SignatureMethod    string
    Timestamp    string
    SignatureVersion   string
    SignatureNonce       string
}


////////////////////////////////////////////////视频上传/////////////////////////////////////////////////////

/**
 * 获取视频上传地址和凭证
 * @link https://help.aliyun.com/document_detail/55407.html?spm=5176.doc44432.6.599.SppqG9
 */
type CreateUploadVideoRequest struct {
    //AlivodRequest
    Action          string              //CreateUploadVideo
    Title           string
    FileName           string
    FileSize         string
    //非必须参数
    Description         string
    CoverURL          string
    CateId       string
    Tags      string
}


/**
 * 获取图片上传地址和凭证
 * @link https://help.aliyun.com/document_detail/55619.html?spm=5176.doc55408.6.605.TATVz1
 */
type CreateUploadImageRequest struct {
    Action          string          //CreateUploadImage
    ImageType           string
    //非必须参数
    ImageExt           string
}


/**
 * 刷新视频上传凭证
 * @link https://help.aliyun.com/document_detail/55408.html?spm=5176.doc55619.6.604.1wj0hR
 */
type RefreshUploadVideoRequest struct {
    Action          string              //RefreshUploadVideo
    VideoId         string
}








////////////////////////////////////////////////视频播放/////////////////////////////////////////////////////

/**
 * 获取视频播放地址
 * @link https://help.aliyun.com/document_detail/52833.html?spm=5176.doc56124.6.609.ZftAQc
 */
type GetPlayInfoRequest struct {
    Action          string              //GetPlayInfo
    VideoId         string
    //非必须参数
    Formats         string
    AuthTimeout     string
}


/**
 * 获取视频播放凭证
 * @link https://help.aliyun.com/document_detail/52833.html?spm=5176.doc56124.6.609.ZftAQc
 */
type GetVideoPlayAuthRequest struct {
    Action          string              //GetVideoPlayAuth
    VideoId         string
}


////////////////////////////////////////////////视频管理/////////////////////////////////////////////////////



/**
 * 获取视频信息
 * @link https://help.aliyun.com/document_detail/52835.html?spm=5176.doc56124.6.611.owFuJq
 */
type GetVideoInfoRequest struct {
    Action          string              //GetVideoInfo
    VideoId         string
}



//等等