package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"os"
)

func Dbupdate(db *sql.DB, arg ...interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("update employees.h_session set timestamp = (?) where open_id = (?)")
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(arg[1], arg[0])
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	err = tx.Commit()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	return err
}

func Dbselect(db *sql.DB, arg ...interface{}) ([]interface{}, error) {
	var open_id string
	var session_id string
	var session_key string
	var timestamp int64
	var temp []interface{}
	stmt, err := db.Prepare("select open_id,session_id,session_key,timestamp from employees.h_session where open_id = (?)")
	if err != nil {
		Srvlog.Error(err.Error())
		return temp, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(arg[0])
	err = row.Scan(&open_id, &session_id, &session_key, &timestamp)
	if err != nil {
		Srvlog.Error(err.Error())
		return temp, err
	}
	temp = append(temp, open_id)
	temp = append(temp, session_id)
	temp = append(temp, session_key)
	temp = append(temp, timestamp)
	return temp, err
}

func Dbinsert(db *sql.DB, arg ...interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO employees.h_session(open_id,session_id,session_key,timestamp) VALUES(?,?,?,?)")
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(arg[0], arg[1], arg[2], arg[3])
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	Srvlog.Info("normal", "ID", lastId, "affected", rowCnt)
	err = tx.Commit()
	if err != nil {
		Srvlog.Error(err.Error())
		//os.Exit(1)
		return err
	}
	return nil
}

func DbselectHuser(db *sql.DB, arg ...interface{}) error {
	var openID string
	stmt, err := db.Prepare("select user_id from employees.h_user where user_id = (?)")
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(arg[0])
	err = row.Scan(&openID)
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	return nil
}

func DbinsertHuser(db *sql.DB, arg ...interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO employees.h_user(user_id,user_pic,nickname) VALUES(?,?,?)")
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(arg[0], arg[1], arg[2])
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	Srvlog.Info("normal", "ID", lastId, "affected", rowCnt)
	err = tx.Commit()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	return nil
}

func DbupdateHuser(db *sql.DB, arg ...interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("update employees.h_user set user_pic = (?),nickname = (?) where user_id = (?)")
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(arg[1], arg[2], arg[0])
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	err = tx.Commit()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	return nil
}
