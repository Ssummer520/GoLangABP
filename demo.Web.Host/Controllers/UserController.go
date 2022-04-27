package Controllers

import (
	service "GoLangABP/demo.Application/User"
	. "GoLangABP/demo.Core/Dto"
	. "GoLangABP/demo.Core/Model"
	"github.com/gin-gonic/gin"
)

// UserController User  注入IStartService
type UserController struct {
	Service *service.UserService `inject:""`
}

// AddUserNameHandler
// @Summary  新增用户信息
// @Tags     User 用户相关
// @Accept   application/json
// @Produce  application/json
// @Router   /user [post]
// @Param object body UserAddInputDto false "请求参数"
func (u *UserController) AddUserNameHandler(c *gin.Context) {
	retObject := RetObject{}
	var input = UserAddInputDto{}
	err := c.ShouldBind(&input)
	if err != nil {
		retObject.Success = false
		retObject.Data = false
		retObject.Message = "不合法的输入数据"
	} else {

		isExist := u.Service.CheckUserInfo(input.Name, input.Phone)
		if isExist {
			retObject.Success = false
			retObject.Data = false
			retObject.Message = "姓名或手机号码已存在"
		} else {
			ret := u.Service.Add(input)
			retObject.Data = ret.Success
			retObject.Success = ret.Success
			retObject.Message = ret.Error
		}
	}
	c.JSON(200, retObject)
}

// GetUserListHandler
// @Summary  获取用户列表
// @Tags     User 获取用户列表
// @Accept   application/json
// @Produce  application/json
// @Router   /user/list [get]
//@Success 200 object RetObject{data=[]UserListOutPutDto} 成功后返回值
// @Param Authorization header string false "Bearer 用户令牌"
func (u *UserController) GetUserListHandler(c *gin.Context) {
	c.Abort()
	retObject := RetObject{}
	ret := u.Service.List()

	if ret == nil && len(ret) == 0 {
		retObject.Success = false
		retObject.Message = "未获取到任何数据"
	} else {
		retObject.Data = ret
	}
	c.JSON(200, retObject)
}
