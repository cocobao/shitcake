package controller

var IsLogin bool

type LogController struct {
	BaseController
}

func (c *LogController) Login() {
	c.TplName = "login.html"
}

func (c *LogController) LoginCommit() {
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")

	if uname != "admin" || pwd != "admin" {
		c.Redirect("/login", 301)
		return
	}
	IsLogin = true
	c.Redirect("/upload", 0)
	// c.TplName = "uploadImage.html"
}
