package handles

import (
	"github.com/kataras/iris/v12"
)

// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept mpfd
// @Produce mpfd
// @Param who query string true "人名"
// @Success 200 {string} string "成功请求响应"
// @Failure 400 {string} string "失败请求响应"
// @Router /hello [get]
func SayHello(c iris.Context) {
	who := c.URLParam("who")
	c.WriteString("hello " + who)
}

type UserInfo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Company string `json:"company"`
}

// @Summary 根据用户ID获取用户信息
// @Description get userInfo
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param id query int true "用户id"
// @Success 200 {object} UserInfo "成功请求响应"
// @Failure 400 {object} ErrResponse "失败请求响应"
// @Router /getUser [get]
func PostInfo(c iris.Context) {
	id, _ := c.URLParamInt("id")
	if id == 1 {
		c.JSON(UserInfo{
				ID:      id,
				Name:    "林鹏",
				Age:     25,
				Company: "XXXXX",
		})
		return
	}

	c.JSON(ErrResponse{
		Status: false,
		ErrMsg:   "失败",
	})
}

type ErrResponse struct {
	Status bool `json:"status"`
	ErrMsg string `json:"errMsg"`
}
