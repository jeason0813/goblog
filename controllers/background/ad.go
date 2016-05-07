package background

import (
	"github.com/deepzz0/goblog/models"
)

type ADController struct {
	Common
}

func (this *ADController) Get() {
	this.Data["Title"] = "广告管理 - " + models.Blogger.BlogName
	this.LeftBar("ad")
	this.Content()
}

func (this *ADController) Content() {
	this.Data["Content"] = ""
}
