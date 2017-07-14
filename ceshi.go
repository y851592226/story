package main

import "fmt"
// import "story/utils/mysql"
import "story/models"
import "os"
import "html/template"

// import "github.com/astaxie/beego"
// import "github.com/garyburd/redigo/redis"
// import "time"
// import "reflect"

type Person struct {
	UserName string
}
func main() {
	chapter,err:= models.GetChapter(2876,1)
	if err!=nil{
		panic(err)
	}
	// t := template.New("ceshi.tpl") //创建一个模板
	// t, _ = t.Parse("hello {{.UserName}}!")
	// p := Person{UserName: "Astaxie"}
	// t.Execute(os.Stdout, p)
	// t, err := template.ParseFiles("views/head.tpl","views/ceshi.tpl")  //解析模板文件
	t, err := template.ParseFiles("views/ceshi.tpl")  //解析模板文件
	// t=t.New(`{{define "head"}}{{.}}{{end}}zzzz`)
	t1,err :=template.New(`head.tpl1`).Parse(`{{define "head"}}{{.Chapter_name}}{{end}}zzzz {{template "line" .}}`)
	t,err = t.AddParseTree("head.tpl2",t1.Tree)
	t,err = t.AddParseTree("head.tpl3",t1.Tree)
	// t, err = t.ParseFiles("views/ceshi.tpl")
// 	t,err = t.New("line.tpl1").Parse(`{{define "line1"}}
// {{.Chapter_name}}{{end}}aaaa
// `)
// 	t = t.New("line.tpl2")
// 	t,err = t.New("line.tpl3").Parse(`{{define "line3"}}
// {{.}}{{end}}{{template "line.tpl2"}}
// `)
	fmt.Println(t.Name())

// 	func unescaped (x string) interface{} { return template.HTML(x) }
// // 在模板对象t中注册unescaped
// t = t.Funcs(template.FuncMap{"unescaped": unescaped})

	if err!=nil{
		panic(err)
	}
	fmt.Println(t.DefinedTemplates())
	// fmt.Println(t.Lookup("line1").DefinedTemplates())

	// user := GetUser() //获取当前用户信息
	err=t.Execute(os.Stdout, *chapter)  //执行模板的merger操作
	if err!=nil{
		panic(err)
	}
	// if chapter==nil{
	// 	panic("nil")
	// }
	// fmt.Println(reflect.TypeOf(res))
	// _, err := mysql.DBStory.Exec("INSERT into user(account) values(?)", "hello")
	// for k,v:=range result[0]{
	//     fmt.Printf("%s %s\n",k,v)
	// }
	// result,err=DBStory.Select("SELECT * FROM book WHERE id = ?",3196)
	// if err!=nil{
	//     panic(err)
	// }
	// for k,v:=range result[0]{
	//     fmt.Printf("%s %s\n",k,v)
	// }
	// result,err=DBStory.Select("SELECT * FROM book WHERE id = ?",3196)
	// if err!=nil{
	//     panic(err)
	// }
	// for k,v:=range result[0]{
	//     fmt.Printf("%s %s\n",k,v)
	// }
	// fmt.Println(chapter)
}
adsjajkkkdaskhello
jasldjfaklskl
aaaa
zzzzzzz
