package controller

import (
	"backend/dao"
	"backend/service"
	"backend/util"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User service.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	usi := service.GetUserServiceInstance()
	user := usi.GetUserBasicInfoByName(username)
	if username == user.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := dao.UserBasicInfo{
			Name:     username,
			Password: service.EnCoder(password),
		}
		if usi.InsertUser(&newUser) != true {
			fmt.Println("Insert Fail")
		}
		// 得到用户id
		user := usi.GetUserBasicInfoByName(username)
		userId := user.Id
		token := util.GenerateToken(userId, username)
		//log.Println("注册时返回的token", token)
		//log.Println("注册返回的id: ", user.Id)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Register Success"},
			UserId:   user.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	dftpassword := "123456"
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	// 使用中间件，做token权限校验
	usi := service.GetUserServiceInstance()
	user := usi.GetUserBasicInfoByName(username)
	if username == user.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: vcode},
		})
	} else {
		newUser := dao.UserBasicInfo{
			Name:     username,
			Password: service.EnCoder(dftpassword),
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: vcode},
		})
		if usi.InsertUser(&newUser) != true {
			fmt.Println("Insert Fail")
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Register Fail"},
			})
		}
	}

}

func GenerateCode(c *gin.Context) {
	username := c.Query("username")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	// 使用中间件，做token权限校验
	usi := service.GetUserServiceInstance()
	user := usi.GetUserBasicInfoByName(username)
	if username == user.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: vcode},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func PasswordLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	encoderPassword := service.EnCoder(password)
	//log.Println("encoderPassword:", encoderPassword)
	log.Println("username:", username)
	// 登录逻辑：使用jwt，根据用户信息生成token
	usi := service.GetUserServiceInstance()

	user := usi.GetUserBasicInfoByName(username)
	userId := user.Id
	if encoderPassword == user.Password {
		token := util.GenerateToken(userId, username)
		//log.Println("generate token:", token)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Login Success"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User or Password Error"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	// 使用中间件，做token权限校验
	usi := service.GetUserServiceInstance()
	id, _ := strconv.ParseInt(userId, 10, 64)
	if user, err := usi.GetUserLoginInfoById(id); err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Query Success"},
			User:     user,
		})
	}
}

func UserExistence(c *gin.Context) {
	username := c.Query("username")
	usi := service.GetUserServiceInstance()
	user := usi.GetUserBasicInfoByName(username)
	// log.Println("this is username: ", username)
	// log.Println("this is user.Name: ", user.Name)
	if username == user.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "User exist!"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist!"},
		})
	}
}

func PasswordChange(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	encoderPassword := service.EnCoder(password)
	if err := dao.Db.Table("user").Where("name = ?", username).Update("password", encoderPassword).Error; err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Password change failed"},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Password change Success"},
		})
	}

}

func Test(c *gin.Context) {
	// 通过c.Get()获取userId
	userId, _ := c.Get("userId")
	//fmt.Println("userId", userId)
	c.JSON(http.StatusOK, userId)
}
