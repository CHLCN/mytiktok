package controller

import (
	"backend/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteActionResponse struct {
	Response
}

type GetFavouriteListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
}

// FavoriteAction 赞操作
func FavoriteAction(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	Fni := service.NewLikeServImpInstance()
	reqUserId := c.Query("user_id")
	userId, _ := strconv.ParseInt(reqUserId, 10, 64)
	err := Fni.FavoriteAction(userId, videoId, int32(actionType))
	if err == nil {
		log.Printf("方法like.FavouriteAction(userid, videoId, int32(actiontype) 成功")
		c.JSON(http.StatusOK, FavoriteActionResponse{
			Response{
				StatusCode: 0,
				StatusMsg:  "favourite action success",
			},
		})
	} else {
		log.Printf("方法like.FavouriteAction(userid, videoId, int32(actiontype) 失败：%v", err)
		c.JSON(http.StatusOK, FavoriteActionResponse{
			Response{
				StatusCode: -1,
				StatusMsg:  "favourite action fail",
			},
		})
	}
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	strUserId := c.Query("user_id")
	//likeCnt:=dao.VideoLikedCount()
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	Fni := service.NewLikeServImpInstance()
	// 返回视频列表信息
	videoList, err := Fni.GetLikesList(userId)
	if err == nil {
		log.Printf("方法like.GetFavouriteList(userid) 成功")
		c.JSON(http.StatusOK, GetFavouriteListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "get favouriteList success"}, VideoList: videoList,
		})
	} else {
		log.Printf("方法like.GetFavouriteList(userid) 失败：%v", err)
		c.JSON(http.StatusOK, Response{StatusCode: -1, StatusMsg: "get favouriteList fail "})
	}
}
