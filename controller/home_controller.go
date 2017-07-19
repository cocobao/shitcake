package controller

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	if !IsLogin {
		c.Redirect("/login", 301)
		return
	}
	c.TplName = "home.html"
}
