package apis

import (
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysDictData struct {
	api.Api
}

// GetSysDictDataList
// @Summary 字典数据列表
// @Description 获取JSON
// @Tags 字典数据
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data [get]
// @Security Bearer
func (e SysDictData) GetSysDictDataList(c *gin.Context) {
	s := service.SysDictData{}
	d := &dto.SysDictDataSearch{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	list := make([]models.SysDictData, 0)
	var count int64
	err = s.GetPage(d, &list, &count)
	if err != nil {
		e.Logger.Errorf("GetPage error, %s", err.Error())
		e.Error(http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), d.GetPageIndex(), d.GetPageSize(), "查询成功")
}

// GetSysDictData
// @Summary 通过编码获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param dictCode path int true "字典编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/{dictCode} [get]
// @Security Bearer
func (e SysDictData) GetSysDictData(c *gin.Context) {
	s := service.SysDictData{}
	d := &dto.SysDictDataById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	var object models.SysDictData

	err = s.Get(d, &object)
	if err != nil {
		e.Logger.Warnf("Get error: %s", err.Error())
		e.Error(http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.OK(object, "查看成功")
}

// InsertSysDictData
// @Summary 添加字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictDataControl true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [post]
// @Security Bearer
func (e SysDictData) InsertSysDictData(c *gin.Context) {
	s := service.SysDictData{}
	d := &dto.SysDictDataControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	// 设置创建人
	d.SetCreateBy(user.GetUserId(c))

	err = s.Insert(d)
	if err != nil {
		e.Logger.Errorf("Insert error, %s", err)
		e.Error(http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(d.GetId(), "创建成功")
}

// UpdateSysDictData
// @Summary 修改字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictDataControl true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data/{dictCode} [put]
// @Security Bearer
func (e SysDictData) UpdateSysDictData(c *gin.Context) {
	s := service.SysDictData{}
	d := &dto.SysDictDataControl{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	d.SetUpdateBy(user.GetUserId(c))
	err = s.Update(d)
	if err != nil {
		e.Logger.Errorf("Update error, %s", err)
		e.Error(http.StatusInternalServerError, err, "更新失败")
		return
	}
	e.OK(d.GetId(), "更新成功")
}

// DeleteSysDictData
// @Summary 删除字典数据
// @Description 删除数据
// @Tags 字典数据
// @Param dictCode path int true "dictCode"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/data/{dictCode} [delete]
// @Security Bearer
func (e SysDictData) DeleteSysDictData(c *gin.Context) {
	s := service.SysDictData{}
	d := new(dto.SysDictDataById)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	// 设置编辑人
	d.SetUpdateBy(user.GetUserId(c))

	err = s.Remove(d)
	if err != nil {
		e.Logger.Errorf("Remove error, %s", err)
		e.Error(http.StatusInternalServerError, err, "删除失败")
		return
	}
	e.OK(d.GetId(), "删除成功")
}

// GetSysDictDataAll 数据字典根据key获取 业务页面使用
// @Summary 数据字典根据key获取
// @Description 数据字典根据key获取
// @Tags 字典数据
// @Param dictType query int true "dictType"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict-data/option-select [get]
// @Security Bearer
func (e SysDictData) GetSysDictDataAll(c *gin.Context) {
	s := service.SysDictData{}
	d := &dto.SysDictDataSearch{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(d, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(http.StatusInternalServerError, err, err.Error())
		e.Logger.Error(err)
		return
	}

	list := make([]models.SysDictData, 0)

	err = s.GetAll(d, &list)
	if err != nil {
		e.Logger.Errorf("GetAll error, %s", err.Error())
		e.Error(http.StatusInternalServerError, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}
