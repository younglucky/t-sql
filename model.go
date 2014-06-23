package main

import (
	_ "code.google.com/p/odbc"
	"database/sql"
	"fmt"
	// "time"
)

func SqlServer() *sql.DB {
	db, err := sql.Open("odbc", "driver={SQL Server};SERVER=goelia2014.sqlserver.rds.aliyuncs.com,3433;UID=young;PWD=goelia;DATABASE=weixin_describe")
	if err != nil {
		panic(err)
	}
	return db
}

// func SqlServer() {
// 	db := SqlServerConn()
// 	defer db.Close()
// 	stmt, err := db.Prepare("select top 5 weixin_id from users")
// 	if err != nil {
// 		fmt.Println("conn Query Error", err)
// 		return
// 	}
// 	defer stmt.Close()
// }

type User struct {
	WeixinId  string `json:"weixin_id"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	Mobile    string `json:"mobile"`
	Birthday  string `json:"birthday"`
}
type DiscribeItem struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Title    string `json:"title"`
}
type Discribes struct {
	Id       int    `json:"id"`
	WeixinId string `json:"weixin_id"`
	ItemId   int    `json:"item_id"`
}

func (d *Discribes) FindByWeixinId(weixinId string) []DiscribeItem2 {
	db := SqlServer()
	defer db.Close()
	rows, err := db.Query("SELECT ISNULL(b.title, '') p_title, a.* FROM discribe_item a LEFT JOIN discribe_item b ON a.parent_id = b.id WHERE a.id IN (SELECT  c.id FROM discribes c WHERE c.weixin_id=?) ORDER BY b.title", weixinId)
	if err != nil {
		panic(err)
	}
	items := make([]DiscribeItem2, 0, 0)
	for rows.Next() {
		var d DiscribeItem2
		err = rows.Scan(&d.PTitle, &d.Id, &d.ParentId, &d.Title)
		if err != nil {
			panic(err)
		}
		items = append(items, d)
	}
	fmt.Println("items", items)
	return items
}

func (u *User) Get(weixinId string) error {
	db := SqlServer()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM users WHERE weixin_id=?", weixinId)
	return row.Scan(&u.WeixinId, &u.Gender, &u.FirstName, &u.Mobile, &u.Birthday)

}
func (u *User) Create(weixinId string) {
	db := SqlServer()
	defer db.Close()
	stmt, err := db.Prepare("insert into users values(?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(weixinId, u.Gender, u.FirstName, u.Mobile, u.Birthday)
	if err != nil {
		panic(err)
	}
}
func (u *User) Update() {
	db := SqlServer()
	defer db.Close()
	stmt, err := db.Prepare("update users set gender=?,first_name=?,mobile=?,birthday=? where weixin_id=?")
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(u.Gender, u.FirstName, u.Mobile, u.Birthday, u.WeixinId)
	if err != nil {
		panic(err)
	}
}
func FindWeixinId(weixinId string) bool {
	// conn := SqlServerConn()
	// defer conn.Close()
	// stmt, err := conn.Prepare("select count(1) from users")
	// if err != nil {
	// 	panic(err)
	// }
	// defer stmt.Close()
	if weixinId == "wahaha" || weixinId == "wahaha1" {
		return true
	}
	return false
}

type DiscribeItem2 struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Title    string `json:"title"`
	PTitle   string `json:"p_title"`
}

func (d *DiscribeItem) All() []DiscribeItem2 {
	db := SqlServer()
	defer db.Close()
	rows, err := db.Query("SELECT ISNULL(b.title, '') p_title, a.* FROM discribe_item a LEFT JOIN discribe_item b ON a.parent_id = b.id ORDER BY b.title")
	if err != nil {
		panic(err)
	}
	items := make([]DiscribeItem2, 0, 0)
	for rows.Next() {
		var d DiscribeItem2
		err = rows.Scan(&d.PTitle, &d.Id, &d.ParentId, &d.Title)
		if err != nil {
			panic(err)
		}
		items = append(items, d)
	}
	fmt.Println("items", items)
	return items
}
func (d *DiscribeItem) Create() {
	db := SqlServer()
	defer db.Close()
	stmt, err := db.Prepare("insert into discribe_item(parent_id,title) values(?,?)")
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(d.ParentId, d.Title)
	if err != nil {
		panic(err)
	}
}
func (d *DiscribeItem) Delete(id string) {
	db := SqlServer()
	defer db.Close()
	stmt, err := db.Prepare("delete discribe_item  where id=?")
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}

func (u *User) All() []User {
	db := SqlServer()
	defer db.Close()
	rows, err := db.Query("SELECT * from users")
	if err != nil {
		panic(err)
	}
	users := make([]User, 0, 0)
	for rows.Next() {
		var u User
		err = rows.Scan(&u.WeixinId, &u.Gender, &u.FirstName, &u.Mobile, &u.Birthday)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}
	fmt.Println("users", users)
	return users
}
