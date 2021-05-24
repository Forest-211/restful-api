package model

import (
	"log"
	"restful/database"
)

type UserModel struct {
	Id       int  `form:"id"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

// CRUD操作
func (user UserModel) Insert() int {
	result, e := database.Db.Exec("insert into `user` (email, password) values (?, ?)", user.Email, user.Password)
	if e != nil {
		log.Panicln("缺少数据！")
	}

	i, _ := result.LastInsertId()
	// 返回插入的id
	return int(i)
}

func (user UserModel) FindAll() []UserModel {
	rows, e := database.Db.Query("select * from `user`")
	if e != nil {
		log.Panicln("查询失败")
	}

	var users []UserModel

	for rows.Next() {
		var u UserModel
		if e := rows.Scan(&u.Id, &u.Email, &u.Password); e == nil {
			users = append(users, u)
		}

	}

	return users
}

func (user UserModel) FindById() UserModel {
	row := database.Db.QueryRow("select * from `user` where id=?", user.Id)
	
	if e := row.Scan(&user.Id, &user.Email, &user.Password); e != nil {
		log.Panicln("绑定发生错误！", e.Error())
	}
	return user
}

func (user UserModel) DeleteOne() string {
	_, e := database.Db.Exec("delete from `user` where id=?", user.Id)
	if e != nil{
		log.Panicln("删除错误", e.Error())
	}
	return "remove successfull"
}
