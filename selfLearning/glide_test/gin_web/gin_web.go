package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/bitly/go-simplejson" // for json get
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/inconshreveable/log15"
	"github.com/jmoiron/sqlx"
)

var Srvlog log.Logger
var TimeDuration int64

func init() {
	Srvlog = log.New()
	Srvlog.SetHandler(log.LvlFilterHandler(log.LvlDebug, log.CallerFileHandler(log.StdoutHandler)))
	TimeDuration = 1800000000000
	//TimeDuration = 1

}

func getOpenid(c *gin.Context) {
	type User struct {
		JsCode string `json:"js_code" binding:"required"`
	}
	var user User
	urlstring := "https://api.weixin.qq.com/sns/jscode2session?appid=wx2e5587abe4bb1323&secret=80fdfe7a6c20943309c70a20baf64bd8&js_code="
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		err := c.BindJSON(&user)
		if err != nil {
			Srvlog.Error(err.Error())
			return
		}
	}
	urlstring = urlstring + user.JsCode + "&grant_type=authorization_code"
	Srvlog.Info("url detail", "url", urlstring, "code", user.JsCode)
	resp, err := http.Get(urlstring)
	if err != nil {
		Srvlog.Error("urlstring get err", "urlstring", urlstring, "err", err)
		return
	} else {
		defer resp.Body.Close()
		bodyCli, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Srvlog.Error(err.Error())
			return
		}
		Srvlog.Info("bodyCli", "bodyCli", string(bodyCli))
		js, err := simplejson.NewJson(bodyCli)
		if err != nil {
			Srvlog.Error(err.Error())
			return
		} else {
			js_session_key := js.Get("session_key").MustString() //new
			js_openid := js.Get("openid").MustString()           //new
			js_timestamp := time.Now().Unix()                    //new
			temp := c.MustGet("db")
			db, ok := temp.(*sqlx.DB)
			if !ok {
				Srvlog.Error("can't type assert")
				return
			} else {
				_, err := DbselectHsession(db, js_openid)
				if err != nil {
					if err == sql.ErrNoRows {
						hsessionInsert := &Hsession{
							OpenID:     js_openid,
							SessionID:  sql.NullString{String: js_openid, Valid: true},
							SessionKey: sql.NullString{String: js_session_key, Valid: true},
							TimeStamp:  sql.NullInt64{Int64: js_timestamp, Valid: true},
						}
						err := DbinsertHsession(db, hsessionInsert)
						if err == nil {
							c.JSON(http.StatusOK, gin.H{
								"status":    1000,
								"sessionid": js_openid,
							})
						}

					}
				} else {
					hsessionUpdate := &Hsession{
						OpenID:    js_openid,
						TimeStamp: sql.NullInt64{Int64: js_timestamp, Valid: true},
					}
					err := DbupdateHsession(db, hsessionUpdate)
					if err == nil {
						c.JSON(http.StatusOK, gin.H{
							"status":    1000,
							"sessionid": js_openid,
						})

					}

				}
			}

		}
	}
}

func MiddleWare(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	db, err := sqlx.Open("mysql", "root:my-secret-pw@tcp(172.17.0.1:3306)/employees")
	if err != nil {
		Srvlog.Crit(err.Error())
		os.Exit(1)
		return
	}
	defer db.Close()
	gin.SetMode(gin.DebugMode)
	// Starts a new Gin instance with no middle-ware
	r := gin.New()
	r.Use(MiddleWare(db))
	v1 := r.Group("/v1.0.0")
	{

	}
	// Define your handlers
	r.POST("/login", getOpenid)
	r.Run()

}
