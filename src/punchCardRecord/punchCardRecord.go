package main

// CREATE DATABASE `attendanceRecord`CHARACTER SET utf8 COLLATE utf8_unicode_ci;

// CREATE TABLE `RECORD`(
//    `id` INT(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
//    `employee_id` INT(20) NOT NULL COMMENT '员工ID',
//    `record_tm` INT(20) NOT NULL COMMENT '记录时间',
//    `day` INT(11) NOT NULL COMMENT '本月第几天',
//    PRIMARY KEY (`id`)
//  )  ENGINE=MYISAM COMMENT='打卡记录' ROW_FORMAT=DEFAULT CHARSET=utf8;

// SELECT  r.employee_id,DATE_FORMAT(FROM_UNIXTIME(MIN(r.record_tm)),'%Y-%m-%d') AS '日期',
// 	DAYOFWEEK(FROM_UNIXTIME(MIN(r.record_tm)))-1 AS '星期',
// 	FROM_UNIXTIME(MIN(r.record_tm)) AS '上班时间',
// 	FROM_UNIXTIME(MAX(r.record_tm)) AS '下班时间'
//   FROM RECORD r
//   WHERE r.record_tm >UNIX_TIMESTAMP('2017-03-01') AND r.record_tm<UNIX_TIMESTAMP('2017-03-31 23:59') GROUP BY r.employee_id,r.day;

import (
	"database/sql"
	"fmt"
	_ "github.com/GO-SQL-Driver/MySQL"
	"time"
)

func main() {

	// fmt.Println(time.Now().YearDay())

	db, err := sql.Open("mysql", "root:root@tcp(172.17.80.6:3306)/attendanceRecord?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT `RECORD`(`employee_id`,`record_tm`,`day`) VALUES ( ?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("904", time.Now().Unix(), time.Now().YearDay())
	checkErr(err)

	fmt.Println(res, "success ok")

	db.Close()

	time.Sleep(3 * time.Second)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
}
