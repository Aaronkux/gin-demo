package system

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BranchApi struct{}

func (b *BranchApi) CreateBranch(c *gin.Context) {
	var r systemReq.CreateBranch
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.BranchCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if branchRes, err := branchService.CreateBranch(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysBranchResponse{Branch: branchRes}, "创建成功", c)
	}
}

func (b *BranchApi) GetBranchList(c *gin.Context) {
	var r systemReq.SearchBranchParams
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if branchList, total, err := branchService.GetBranchList(r); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("获取失败", err, c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     branchList,
			Total:    total,
			Page:     r.Page,
			PageSize: r.PageSize,
		}, "", c)
	}
}

func (b *BranchApi) GetBranchById(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	branchRes, err := branchService.GetBranchById(r.ID)
	if err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysBranchResponse{Branch: branchRes}, "", c)
	}
}

func (b *BranchApi) UpdateBranch(c *gin.Context) {
	var r system.SysBranch
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.BranchUpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := branchService.UpdateBranch(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("更新失败", err, c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (b *BranchApi) DeleteBranch(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := branchService.DeleteBranch(r.ID); err != nil {
		global.AM_LOG.Error("删除成功!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("删除成功", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
