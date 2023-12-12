package controller

import (
	"backend/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
	NextTime  int64           `json:"next_time"`
}

// Feed 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
func Feed(c *gin.Context) {
	var convTime time.Time
	convTime = time.Now()
	//userId := c.GetInt64("userId")
	strUserId := c.Query("userId")
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	videoService := service.GetVideoServiceInstance()
	videos, nextTime, err := videoService.Feed(convTime, userId)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1, StatusMsg: "刷新视频流失败"},
			VideoList: nil,
			NextTime:  nextTime.Unix(),
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "刷新视频流成功!"},
		VideoList: videos,
		NextTime:  nextTime.Unix(),
	})
}
