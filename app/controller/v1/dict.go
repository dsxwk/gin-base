package v1

import (
	"gin-base/app/service"
	"gin-base/common"
	"gin-base/common/global"
	"github.com/gin-gonic/gin"
)

type DictController struct {
	common.BaseController
}

// List 列表
// @Router /api/v1/dict [get]
func (s *DictController) List(c *gin.Context) {
	var (
		dictService service.DictService
	)

	data, err := dictService.List()
	if err != nil {
		global.Log.Error(err.Error())
		s.ApiResponse(c, global.SystemError, err.Error())
		return
	}

	s.ApiResponse(c, global.Success, data)
}
