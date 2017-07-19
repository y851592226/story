package controllers

import (
	"github.com/astaxie/beego"
	"story/models"
	"strconv"
)

type ChapterController struct {
	beego.Controller
}

type BookInfoController struct {
	beego.Controller
}

func (c *ChapterController) Get() {
	book_id, _ := strconv.Atoi(c.Ctx.Input.Param(":book_id"))
	chapter_id, _ := strconv.Atoi(c.Ctx.Input.Param(":chapter_id"))
	chapter, err := models.GetChapter(book_id, chapter_id)
	if err != nil {
		c.Abort("500")
	} else if chapter == nil {
		c.Abort("404")
	} else {
        c.Data["topbar"] = models.GetTopbar2(chapter.Book_name)
        c.Data["subnav"] = models.GetSubnav()
		c.Data["chapter"] = chapter
		c.Data["Title"] = chapter.Chapter_name
		c.TplName = "chapter.tpl"
	}
}

func (c *BookInfoController) Get() {
	// book_id := c.Ctx.Input.Param(":book_id")
	// book = models.GetBook(book_id)
	// for k,v := range chapter{
	//     c.Data[k] = v
	// }
	// c.TplName = "book.html"
}
