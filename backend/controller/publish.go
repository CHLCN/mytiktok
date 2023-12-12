package controller

import (
	"backend/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	userId := c.Query("author_id")
	log.Println("userId: ", userId)
	autherId, _ := strconv.ParseInt(userId, 10, 64)
	title := c.Query("title")
	playUrl := c.Query("play_url")
	coverUrl := c.Query("cover_url")
	videoService := service.GetVideoServiceInstance()
	err := videoService.Publish(autherId, title, playUrl, coverUrl)
	if err != nil {
		log.Println("上传文件失败")
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  fmt.Sprintf("《%s》视频上传成功", title),
	})
}

// PublishList 用户的视频发布列表，直接列出用户所有投稿过的视频
func PublishList(c *gin.Context) {
	// 获取到的userId，被访问的用户id
	reqUserId := c.Query("user_id")
	userId, _ := strconv.ParseInt(reqUserId, 10, 64)
	//userId := c.GetInt64("userId")
	log.Println("获取到用户 Id：", userId)
	token := c.Query("token")
	log.Println("获取到用户 token：", token)
	videoService := service.GetVideoServiceInstance()
	videos, err := videoService.PublishList(userId)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1, StatusMsg: "获取用户视频发布列表失败!"},
			VideoList: nil,
		})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取用户发布的视频列表成功！",
		},
		VideoList: videos,
	})
}
