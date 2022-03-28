package request

type FileAuthorize struct {
	FileName  string `form:"filename"`
	Code      int    `form:"code"`
	AuthToken string `form:"auth_token"`
	Scene     string `form:"scene"`
	Path      string `form:"path"`
}
