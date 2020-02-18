package main

import (
	"database/sql"
	//"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Hsession struct {
	OpenID     string         `db:"open_id"`
	SessionID  sql.NullString `db:"session_id"`
	SessionKey sql.NullString `db:"session_key"`
	TimeStamp  sql.NullInt64  `db:"timestamp"`
}

type Huser struct {
	UserID      string          `db:"user_id"`
	Birth       sql.NullString  `db:"birth"`
	CreateDate  sql.NullInt64   `db:"create_data"`
	Gender      sql.NullInt64   `db:"gender"`
	Lat         sql.NullFloat64 `db:"lat"`
	Lon         sql.NullFloat64 `db:"lon"`
	Nickname    sql.NullString  `db:"nickname"`
	UserPic     sql.NullString  `db:"user_pic"`
	UserType    sql.NullInt64   `db:"user_type"`
	AudtiStatus sql.NullInt64   `db:"audti_status"`
}

func DbselectHsession(db *sqlx.DB, openID string) (Hsession, error) {
	var session Hsession
	err := db.Get(&session, "select * from employees.h_session where open_id = (?)", openID)
	if err != nil {
		Srvlog.Error(err.Error())
		return session, err
	}
	return session, nil
}

func DbinsertHsession(db *sqlx.DB, session *Hsession) (result error) {
	tx := db.MustBegin()
	defer tx.Rollback()
	_, err := tx.NamedExec("INSERT IN employees.h_session(open_id,session_id,session_key,timestamp) VALUES(:open_id,:session_id,:session_key,:timestamp)", session)
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	//tx.MustExec("INSERT IN employees.h_session(open_id,session_id,session_key,timestamp) VALUES(?,?,?,?)", hsessionInsert.open_id, hsessionInsert.session_id, hsessionInsert.session_key, hsessionInsert.timestamp")
	err = tx.Commit()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	return err
}

func DbupdateHsession(db *sqlx.DB, session *Hsession) error {
	tx := db.MustBegin()
	defer tx.Rollback()
	//tx.MustExec("update employees.h_session set timestamp = (?) where open_id = (?)", timestamp, openID)
	_, err := tx.NamedExec("update employees.h_session set timestamp=(:timestamp) where open_id = (:open_id)", session)
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

func DbselectHuser(db *sqlx.DB, userID string) (Huser, error) {
	var user Huser
	err := db.Get(&user, "select * from employees.h_user where user_id = (?)", userID)
	if err != nil {
		Srvlog.Error(err.Error())
		return user, err
	}
	return user, err
}

func DbinsertHuser(db *sqlx.DB, user *Huser) (result error) {
	tx := db.MustBegin()
	defer tx.Rollback()
	_, err := tx.NamedExec("INSERT IN employees.h_user(user_id,birth,create_date,gender,lat,lon,nickname,user_pic,user_type,audi_status) VALUES(:user_id,:birth,:create_date,:gender,:lat,:lon,:nickname,:user_pic,:user_type,:audi_status)", user)
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	//tx.MustExec("INSERT IN employees.h_session(open_id,session_id,session_key,timestamp) VALUES(?,?,?,?)", hsessionInsert.open_id, hsessionInsert.session_id, hsessionInsert.session_key, hsessionInsert.timestamp")
	err = tx.Commit()
	if err != nil {
		Srvlog.Error(err.Error())
		return err
	}
	return err
}

func DbupdateHuser(db *sqlx.DB, user *Huser) error {
	tx := db.MustBegin()
	defer tx.Rollback()
	//tx.MustExec("update employees.h_session set timestamp = (?) where open_id = (?)", timestamp, openID)
	_, err := tx.NamedExec("update employees.h_user set birth=(:birth),create_date=(:create_date),gender=(:gender),lat=(:lat),lon=(:lon),nickname=(:nickname),user_pic=(:user_pic),uesr_type=(:user_type),audti_status=(:audti_status) where user_id = (:user_id)", user)
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
