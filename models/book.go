package models

import (
	"html/template"
	"story/utils/mysql"
	"strconv"
	"strings"
	"fmt"
)

type Item struct {
	Href template.URL
	Name string
}

type Topbar struct {
	Item1 *Item
	Item2 *Item
	Item3 *Item
}

type Subnav struct {
	Item []*Item
}

func GetSubnav() (subnav Subnav) {
	item1 := &Item{Href: "/", Name: "首页"}
	item2 := &Item{Href: "/classify", Name: "分类"}
	item3 := &Item{Href: "/rank", Name: "排行"}
	item4 := &Item{Href: "/sf", Name: "轻小说"}
	item5 := &Item{Href: "/shujia", Name: "书架"}
	subnav.Item = append(subnav.Item, item1, item2, item3, item4, item5)
	return
}
func GetTopbar2(book_name string) (topbar Topbar) {
	item1 := &Item{Href: "javascript:history.go(-1);", Name: "返回"}
	item2 := &Item{Href: "", Name: book_name}
	item3 := &Item{Href: "/", Name: "首页"}
	topbar.Item1 = item1
	topbar.Item2 = item2
	topbar.Item3 = item3
	return
}

type Chapter struct {
	Book_id      int      `json:"book_id"`
	Book_name    string   `json:"book_name"`
	Chapter_num  int      `json:"chapter_num"`
	Chapter_name string   `json:"chapter_name"`
	Chapter_url  string   `json:"chapter_url"`
	Href_mulu    string   `json:"href_mulu"`
	Href_before  string   `json:"href_before"`
	Href_after   string   `json:"href_after"`
	Content      []string `json:"content"`
}

type CChapter struct {
	Chapter_num  int    `json:"chapter_num"`
	Chapter_name string `json:"chapter_name"`
	Read_flag    bool   `json:"read_flag"`
}

type Catalog struct {
	Chapters []*CChapter `json:"chapters"`
}

type Book struct {
	Book_id           int      `json:"book_id"`
	Book_name         string   `json:"book_name"`
	Book_url          string   `json:"book_url"`
	Book_author       string   `json:"book_author"`
	Book_intro        string   `json:"book_intro"`
	Book_type1        string   `json:"book_type1"`
	Book_type2        string   `json:"book_type2"`
	Book_img          string   `json:"book_img"`
	Last_read_chapter *Chapter `json:"last_read_chapter"`
	Latest_chapter    *Chapter `json:"latest_chapter"`
	Catalog           *Catalog `json:"catalog"`
}

func getCatalogFromMysql(book_id int) (catalog *Catalog, err error) {
	sql := `select chapter_name,chapter_num from chapter_detail  where book_id= ?
        order by chapter_num`
	result, err := mysql.DBStory.Select(sql, book_id)
	if err != nil {
		return
	}
	if len(result) == 0 {
		return
	}
    catalog = new(Catalog)
	for _, item := range result {
		chapter := new(CChapter)
		chapter.Chapter_name = string(item[`chapter_name`].([]byte))
		chapter.Chapter_num = int(item[`chapter_num`].(int64))
		catalog.Chapters = append(catalog.Chapters, chapter)
	}
	return
}
func getBookFromMysql(book_id int) (book *Book, err error) {
    // sql := `SELECT chapter_name,content,book_name,chapter_num,book_id 
    //     FROM chapter_detail where book_id = ? and chapter_num = ?`
    // result, err := mysql.DBStory.Select(sql, book_id, 3)
	sql := `SELECT id,book_name,book_url,book_author,book_intro,book_type1,
        book_type2,book_img from book where id= ?`
	result, err := mysql.DBStory.Select(sql, book_id)
	if err != nil {
		return
	}
    fmt.Println(result,book_id)
	if len(result) == 0 {
		return
	}
    book = new(Book)
	book.Book_id = int(result[0][`id`].(int64))
	book.Book_name = string(result[0][`book_name`].([]byte))
	book.Book_url = string(result[0][`book_url`].([]byte))
	book.Book_author = string(result[0][`book_author`].([]byte))
	book.Book_intro = string(result[0][`book_intro`].([]byte))
	book.Book_type1 = string(result[0][`book_type1`].([]byte))
	book.Book_type2 = string(result[0][`book_type2`].([]byte))
	book.Book_img = string(result[0][`book_img`].([]byte))
	// book.Last_read_chapter=
	// book.Latest_chapter=
	Catalog, err := getCatalogFromMysql(book_id)
	if err != nil {
		return
	}
	book.Catalog = Catalog
	return
}

func getBookFromRedis(book_id int) (book *Book, err error) {
	return
}

func getBookFromSpider(book_id int) (book *Book, err error) {
	return
}

func getChapterFromMysql(book_id, chapter_num int) (chapter *Chapter, err error) {
	sql := `SELECT chapter_name,content,book_name,chapter_num,book_id 
        FROM chapter_detail where book_id = ? and chapter_num = ?`
	result, err := mysql.DBStory.Select(sql, book_id, chapter_num)
	if err != nil {
		return
	}
	if len(result) == 0 {
		return
	}
	chapter = new(Chapter)
	// fmt.Println(result[0][`book_id`].(int64))
	chapter.Book_id = int(result[0][`book_id`].(int64))
	chapter.Book_name = string(result[0][`book_name`].([]byte))
	chapter.Chapter_num = int(result[0][`chapter_num`].(int64))
	chapter.Chapter_name = string(result[0][`chapter_name`].([]byte))
	chapter.Chapter_url = `/get/book/` + strconv.Itoa(book_id) +
		`/chapter/` + strconv.Itoa(chapter_num)
	chapter.Href_before = chapter.Chapter_url
	chapter.Href_after = chapter.Chapter_url
	chapter.Href_mulu = chapter.Chapter_url
	chapter.Content = strings.Split(strings.Replace(string(result[0][`content`].([]byte)), " ", "&nbsp;", -1), "\r\n")
	return
}

func getChapterFromRedis(book_id, chapter_num int) (chapter *Chapter, err error) {
	return

}

func getChapterFromSpider(book_id, chapter_num int) (chapter *Chapter, err error) {
	return

}

func GetChapter(book_id, chapter_num int) (chapter *Chapter, err error) {
	chapter, err = getChapterFromMysql(book_id, chapter_num)
	return
}

func GetBook(book_id int) (book *Book, err error) {
	book, err = getBookFromMysql(book_id)
	return

}
