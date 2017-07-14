package controllers

import (
    "github.com/astaxie/beego"
    // "story/models"
    // "strconv"
)

type ChapterController struct {
    beego.Controller
}

type BookInfoController struct {
    beego.Controller
}

func (c *ChapterController) Get() {
    // book_id,_ := strconv.Atoi(c.Ctx.Input.Param(":book_id"))
    // chapter_id,_ := strconv.Atoi(c.Ctx.Input.Param(":chapter_id"))
    // chapter,_ := models.GetChapter(book_id,chapter_id)
    // // for k,v := range *chapter{
    // //     c.Data[k] = v
    // // }
    // c.Data["Title"] = c.Data["Chapter_name"]
    // c.TplName = "chapter.html"
}

func (c *BookInfoController) Get() {
    // book_id := c.Ctx.Input.Param(":book_id")
    // book = models.GetBook(book_id)
    // for k,v := range chapter{
    //     c.Data[k] = v
    // }
    // c.TplName = "book.html"
}