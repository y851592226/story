package controllers

import (
	"github.com/astaxie/beego"
	"story/models"
	"strconv"
    "fmt"
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
	book_id, _ := strconv.Atoi(c.Ctx.Input.Param(":book_id"))
    fmt.Println(book_id)
	book, err := models.GetBook(book_id)
	if err != nil {
        // panic(err)
		c.Abort("500")
	} else if book == nil {
        // panic("book not found")
		c.Abort("404")
	} else {
		c.Data["topbar"] = models.GetTopbar2(book.Book_name)
		c.Data["subnav"] = models.GetSubnav()
		c.Data["book"] = book
		c.Data["Title"] = book.Book_name
		c.TplName = "book.tpl"
	}
}
