package service

import (
	"github.com/weilaim/blog-api/model"
	"github.com/weilaim/blog-api/serializer"
)

// ShowVideoService 视频详情服务
type ShowVideoService struct {
	 
}

func (service *ShowVideoService)Show(id string) serializer.Response{
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}