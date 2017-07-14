package controllers

import (
    "github.com/astaxie/beego"
    "story/utils/mysql"
    "strings"
    // "time"
    // "fmt"
)
type response struct {
    Err_code int `json:"err_code"`
    Message string `json:"message"`
}
type LoginController struct {
    beego.Controller
}

type RegisterController struct {
    beego.Controller
}

type LogoutController struct {
    beego.Controller
}

func Login(username, password string) (res response, err error){
    sql := `select password from user where username=?`
    result, err := mysql.DBStory.Select(sql,username)
    if err != nil{
        res.Err_code = -1
        res.Message = `服务异常`
        err = nil
        return
    }
    if len(result) == 0{
        res.Err_code = -2
        res.Message = `用户名不存在`
        return
    }
    if string(result[0][`password`].([]byte)) != password{
        res.Err_code = -3
        res.Message = `密码错误`
        return
    }
    res.Err_code = 0
    res.Message = `登录成功`
    return 
}

func Logout(c *LogoutController) (res response, err error){
    c.Ctx.SetCookie(`user`,`nil`,-1)
    res.Err_code = 0
    res.Message = `登出成功`
    return     
}

func Register(username, password, email string) (res response, err error){
    sql := `insert into user(username,password,email) values(?,?,?)`
    _, err = mysql.DBStory.Exec(sql,username,password,email)
    if err != nil{       
        if strings.Contains(err.Error(),`Duplicate entry`){
            res.Err_code = -2
            res.Message = `该账号已被注册`
            err = nil
        }else{
            res.Err_code = -1
            res.Message = `服务异常`
        }
        return
    }
    res.Err_code = 0
    res.Message = `恭喜你，注册成功！！`
    return
}

func (c *LoginController) Get() {
    c.TplName = "login.html"
}

func (c *LoginController) Post() {
    username := c.GetString("username")
    password := c.GetString("password")
    result, err := Login(username, password)
    if err!=nil{
        c.Abort("500")
    }
    c.Ctx.SetCookie("user",username,30*24*60*60)
    c.Data["json"] = &result
    c.ServeJSON()
}

func (c *RegisterController) Get() {
    c.TplName = "register.html"
}

func (c *RegisterController) Post() {
    username := c.GetString("username")
    password := c.GetString("password")
    email := c.GetString("email")
    result, err := Register(username, password, email)
    if err!=nil{
        c.Abort("500")
    }
    c.Data["json"] = &result
    c.ServeJSON()
}

func (c *LogoutController) Post() {
    result,err := Logout(c)
    if err!=nil{
        c.Abort("500")
    }
    c.Data["json"] = &result
    c.ServeJSON()
}

func (c *LogoutController) Get() {
    _,err := Logout(c)
    if err!=nil{
        c.Abort("500")
    }
    c.Redirect("/",302)
}