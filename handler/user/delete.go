package user

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
