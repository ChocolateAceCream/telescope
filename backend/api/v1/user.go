package apiV1

import (
	"encoding/json"

	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u *UserApi) GetUserInfo(c *gin.Context) {
	user, err := utils.GetValueFromSessionByKey[map[string]interface{}](c, "user")

	// 🔹 Convert map to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		response.FailWithMessage(c, "error.failed.operation")
		return
	}

	// 🔹 Convert JSON to struct
	var userStruct response.GetUserResp
	err = json.Unmarshal(jsonData, &userStruct)
	if err != nil {
		response.FailWithMessage(c, "error.failed.operation")
		return
	}
	response.OkWithFullDetails(c, userStruct, "success")
}

func (u *UserApi) Logout(c *gin.Context) {
	err := utils.DeleteSession(c, singleton.Config.Session.CookieName)
	if err != nil {
		response.FailWithMessage(c, "error.failed.operation")
		return
	}

	response.OkWithMessage(c, "success")
}
