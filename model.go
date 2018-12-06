package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//https://www.cnblogs.com/wt645631686/p/9691606.html
type DB struct {
	Handle    *sql.DB
	Prefix    string
	TableName string
	Condition string
}

const (
	username string = "eatojoy"
	password string = "8Lect0J8kfxDIrXR"
	host     string = "10.12.0.24"
	port     int    = 3306
	dbname   string = "eatojoy"
)

/*
条件：where
方法：getOne/getList/count/sum/update/delete/insert
*/
func (db *DB) DBConnect() *sql.DB {
	link := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, dbname)
	handle, err := sql.Open("mysql", link)

	if err != nil {
		log.Fatal(err)
	}

	return handle
}

func (db *DB) Table(tableName string) *DB {
	db.TableName = tableName
	return db
}

func (db *DB) Where(where string) *DB {
	if db.Condition == "" {
		db.Condition = where
	} else {
		db.Condition += " and " + where
	}

	return db
}

func (db *DB) GetOne(field string) (record map[string]string, err error) {
	db.TableName = db.Prefix + db.TableName
	fmt.Println(db.TableName)

	queryString := fmt.Sprintf(
		"select %s from %s where %s",
		field, db.TableName, db.Condition)

	rows, _ := db.DBConnect().Query(queryString)
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		rows.Scan(scanArgs...)
		record = make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		break
	}
	return
}

func (db *DB) GetList(field string, order string, start int, limit int) (data []map[string]string, err error) {
	db.TableName = db.Prefix + db.TableName
	fmt.Println(db.TableName)

	queryString := fmt.Sprintf(
		"select %s from %s where %s order by %s limit %d,%d",
		field, db.TableName, db.Condition, order, start, limit)

	rows, _ := db.DBConnect().Query(queryString)
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}

		data = append(data, record)
	}
	return
}

func (db *DB) Update(map[string]string) (id int64, err error) {
	stmt, err := db.DBConnect().Prepare(`UPDATE student SET age=? WHERE id=?`)
	res, err := stmt.Exec(21, 5)
	num, err := res.RowsAffected() //影响行数
	fmt.Println(num)
	return
}

func (db *DB) Delete([]string) (id int64, err error) {
	stmt, err := db.DBConnect().Prepare(`DELETE FROM student WHERE id=?`)
	res, err := stmt.Exec(5)
	num, err := res.RowsAffected()
	fmt.Println(num)
	return
}

func (db *DB) insert([]string) (id int64, err error) {
	stmt, err := db.DBConnect().Prepare(`INSERT student (name,age) values (?,?)`)
	res, err := stmt.Exec(21, 5)
	id, err = res.LastInsertId()
	fmt.Println(id)
	return
}
