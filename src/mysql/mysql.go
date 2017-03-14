package main

// CREATE TABLE `userinfo`(
//    `uid` INT(10) NOT NULL AUTO_INCREMENT ,
//    `username` VARCHAR(64) NULL DEFAULT NULL ,
//    `departname` VARCHAR(64) NULL DEFAULT NULL,
//    `created` DATE NULL DEFAULT NULL,
//    PRIMARY KEY (`uid`)
//  );

// CREATE TABLE `userdetail`(
//    `uid` INT(10) NOT NULL DEFAULT '0' ,
//    `intro` TEXT NULL ,
//    `profile` TEXT NULL,
//    PRIMARY KEY (`uid`)
//  );
//
// INSERT INTO `userinfo`(`uid`,`username`,`departname`) VALUES ( '1','fan','研发');

import (
	"database/sql"
	"fmt"
	_ "github.com/GO-SQL-Driver/MySQL"
	// "time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)

	// //插入数据
	// stmt, err := db.Prepare("INSERT `userinfo`(`username`,`departname`,`created`) VALUES ( ?,?,?)")
	// checkErr(err)

	// res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)

	id := 4

	// //更新数据
	// stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err := stmt.Exec("astaxieupdate", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// // 查询数据
	// rows, err := db.Query("select * from userinfo")
	// checkErr(err)

	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var department string
	// 	var created string
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }

	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
