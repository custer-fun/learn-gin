package classes

import (
	"fmt"
	"gin-up/src/goft"
	"gin-up/src/models"
	"github.com/gin-gonic/gin"
)

// UserClass 
type UserClass struct {
	//*goft.GormAdapter
	*goft.XormAdapter
	Age *goft.Value `prefix:"user.age"`
}

// NewUserClass UserClass generate constructor
func NewUserClass() *UserClass {
	return &UserClass{}
}

// UserTest 控制器方法
func (this *UserClass) UserTest(c *gin.Context) string {
	return "用户测试" + this.Age.String()
}

// UserList 用户列表 返回切片
func (this *UserClass) UserList(c *gin.Context) goft.Models {
	users := []*models.UserModel{
		&models.UserModel{UserID: 101, UserName: "custer"},
		{UserID: 102, UserName: "张三"},
		{UserID: 103, UserName: "李四"},
	}
	return goft.MakeModels(users)
}

// UserDetail 返回 Model 实体(即所有自定义的 struct)，返回值都是 goft.Model
func (this *UserClass) UserDetail(c *gin.Context) goft.Model {
	user := models.NewUserModel()
	err := c.BindUri(user)
	goft.Error(err, "ID 参数不合法") // 如果出错会发生 panic，然后在中间件中处理返回 400
	//this.Table(user.TableName()).
	//	Where("user_id=?", user.UserID).Find(user)
	has, err := this.Table("users").
		Where("user_id=?", user.UserID).Get(user)
	if err != nil {
		goft.Error(err)
	}
	if !has { // 没有记录
		goft.Error(fmt.Errorf("没有该用户"))
	}
	return user
}

func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/test", this.UserTest).
		Handle("GET", "/userlist", this.UserList).
		Handle("GET", "/user/:id", this.UserDetail)
}

func (this *UserClass) Name() string {
	return "UserClass"
}
