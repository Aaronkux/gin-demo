package system

import (
	"net/http"
	"path/filepath"

	systemReq "gandi.icu/demo/model/system/request"
	"github.com/gin-gonic/gin"
)

type FileApi struct{}

func (f *FileApi) Authorize(c *gin.Context) {
	var r systemReq.FileAuthorize
	_ = c.ShouldBind(&r)
	println(r.Code, r.AuthToken, r.FileName, r.Scene, filepath.Clean(r.Path))
	c.String(http.StatusOK, "ok")
}

// path=%2Ftest3&submit=upload&scene=globalpay3&__path__=%2Fgroup1%2Fupload&__query__=&code=&auth_token=&filename=&output=json2
