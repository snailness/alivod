# go alivod
aliyun 视频点播 SDK  alivod

## Installation

    go get github.com/snailness/alivod

## Useage
    //init aliyun access_key. you need init only one time.
    alivod.InitConfig("access_key_id", "access_key_secret") 
    // eq. GetVideoInfo
    query_url := alivod.GetAlivodVideoInfoUrl(alivod.GetVideoInfoRequest{
        Action: alivod.ALIVOD_API_GET_VIDEO_INFO,
        VideoId: vid,
    })
    response, err := alivod.GetHttpResponse(query_url)
    if err != nil {
        return nil
    }