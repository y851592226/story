package routers

import (
	"story/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/logout", &controllers.LogoutController{})
    beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/get/book/:book_id:int", &controllers.BookInfoController{})
    beego.Router("/get/book/:book_id:int/chapter/:chapter_id:int", &controllers.ChapterController{})
    // beego.Router("/search/bookname/:name:string", &controllers.SearchBookNameController{})
    // beego.Router("/search/authorname/:name:string", &controllers.SearchBookNameController{})
    // beego.Router("/search/auto/:name:string", &controllers.SearchBookNameController{})
    // beego.Router("/get/classify", &controllers.BookInfoController{})
    // beego.Router("/get/classify/:name:string", &controllers.BookInfoController{})
    // beego.Router("/get/rank", &controllers.BookInfoController{})
    // beego.Router("/get/rank/:name:string", &controllers.BookInfoController{})
    // beego.Router("/get/history", &controllers.BookInfoController{})
    // beego.Router("/get/sanjiang", &controllers.BookInfoController{})
    // beego.Router("/get/strongrec", &controllers.BookInfoController{})
    beego.SetStaticPath("/file", "static")
}
