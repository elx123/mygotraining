package main

import (
	"context"
	//"crdu"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/inconshreveable/log15"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var Srvlog log.Logger
var TimeDuration int64

func init() {
	Srvlog = log.New()
	Srvlog.SetHandler(log.LvlFilterHandler(log.LvlDebug, log.CallerFileHandler(log.StdoutHandler)))
	TimeDuration = 1800000000000
	//TimeDuration = 1

}

func getOpenid(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	urlstring := "https://api.weixin.qq.com/sns/jscode2session?appid=wx2e5587abe4bb1323&secret=80fdfe7a6c20943309c70a20baf64bd8&js_code="
	if r.Method == "POST" {
		if len > 0 {
			body := make([]byte, len)
			r.Body.Read(body)
			js, err := simplejson.NewJson(body)
			if err != nil {
				Srvlog.Error("NewJson err", "err", err)
				return
			}
			js_code_temp := js.Get("js_code").MustString()

			urlstring = urlstring + js_code_temp + "&grant_type=authorization_code"
			Srvlog.Info("url detail", "url", urlstring, "code", js_code_temp)
			resp, err := http.Get(urlstring)
			if err != nil {
				Srvlog.Error("urlstring get err", "urlstring", urlstring, "err", err)
				return
			} else {
				defer resp.Body.Close()
				bodyCli, _ := ioutil.ReadAll(resp.Body)
				Srvlog.Info("bodyCli", "bodyCli", string(bodyCli))
				js, err = simplejson.NewJson(bodyCli)
				if err != nil {
					Srvlog.Error("NewJson err", "err", err)
					return
				} else {
					js_session_key := js.Get("session_key").MustString() //new
					js_openid := js.Get("openid").MustString()           //new
					js_timestamp := time.Now().Unix()                    //new
					temp := r.Context().Value("db")
					db, ok := temp.(*sql.DB)
					if !ok {
						Srvlog.Error("type assert", "err", ok)
						return
					} else {
						_, err := Dbselect(db, js_openid)
						if err != nil {
							if err == sql.ErrNoRows {
								err := Dbinsert(db, js_openid, js_session_key, js_openid, js_timestamp)
								if err == nil {
									js_response := simplejson.New()
									js_response.Set("status", 1000)
									js_response.Set("sessionid", js_openid)
									body, err := js_response.Encode()
									if err != nil {
										Srvlog.Error("simplejson MarshalJson", "err", err)
									}
									w.Write(body)
								}

							}
						} else {
							err := Dbupdate(db, js_openid, js_timestamp)
							if err == nil {
								js_response := simplejson.New()
								js_response.Set("status", 1000)
								js_response.Set("sessionid", js_openid)
								body, err := js_response.Encode()
								if err != nil {
									Srvlog.Error("simplejson MarshalJson", "err", err)
								}
								w.Write(body)
							}

						}
					}
				}
			}
		}
	}
}

func AddContextSupport(next http.Handler, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "db", db)
		// WithContext returns a shallow copy of r with its context changed
		// to ctx. The provided ctx must be non-nil.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, -1)
	cookie := http.Cookie{Name: "username", Value: "alice_cooper@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	if username := r.Context().Value("username"); username != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hi username:" + username.(string) + "\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}
func patter_login(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/mengyou/login" {
			getOpenid(w, r)
		} else {
			h.ServeHTTP(w, r)
		}

	}
}

func patter_session(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("sessionid") == "" {
			w.WriteHeader(http.StatusOK)
			js_response := simplejson.New()
			js_response.Set("status", 4003)
			js_response.Set("error", "session empty")
			body, err := js_response.Encode()
			if err != nil {
				Srvlog.Error("simplejson MarshalJson", "err", err)
			} else {
				Srvlog.Error("Dbselect error", "err", err)
			}
			w.Write(body)
			return
		} else {
			sessionid := r.URL.Query().Get("sessionid")
			js_timestamp := time.Now().Unix()
			temp := r.Context().Value("db")
			db, ok := temp.(*sql.DB)
			if !ok {
				Srvlog.Error("type assert", "err", ok)
				return
			} else {
				result, err := Dbselect(db, sessionid)
				if err != nil {
					w.WriteHeader(http.StatusOK)
					js_response := simplejson.New()
					js_response.Set("status", 4003)
					js_response.Set("error", err.Error())
					body, err := js_response.Encode()
					if err != nil {
						Srvlog.Error("simplejson MarshalJson", "err", err)
					} else {
						Srvlog.Error("Dbselect error", "err", err)
					}
					w.Write(body)
					return
				} else {
					v, _ := result[3].(int64)
					if js_timestamp-v >= TimeDuration {
						js_response := simplejson.New()
						js_response.Set("status", 4003)
						js_response.Set("error", "session_id expired")
						body, err := js_response.Encode()
						if err != nil {
							Srvlog.Error("simplejson MarshalJson", "err", err)
						} else {
							Srvlog.Info("session expired", "openid", sessionid)
						}
						w.Write(body)
						return
					} else {
						Dbupdate(db, sessionid, js_timestamp)
						h.ServeHTTP(w, r)
					}
				}
			}
		}
	}

}

func Test(w http.ResponseWriter, r *http.Request) {
}

func SaveUserinfo(w http.ResponseWriter, r *http.Request) {
	sessionid := r.URL.Query().Get("sessionid")
	len := r.ContentLength
	if r.Method == "POST" {
		if len > 0 {
			body := make([]byte, len)
			r.Body.Read(body)
			js, err := simplejson.NewJson(body)
			if err != nil {
				Srvlog.Error("NewJson err", "err", err)
				return
			}
			userPic := js.Get("userPic").MustString()
			nickname := js.Get("nickname").MustString()
			Srvlog.Info("thing", "userPic", userPic)
			if userPic == "userPic" || nickname == "nickname" {
				js_response := simplejson.New()
				js_response.Set("status", 4003)
				js_response.Set("error", "userPic or nickname is empity")
				body, err := js_response.Encode()
				if err != nil {
					Srvlog.Error("simplejson MarshalJson", "err", err)
				} else {
					Srvlog.Info("userPic or nickname is empity", "err", "userPic or nickname is empity")
				}
				w.Write(body)
				return
			} else {
				temp := r.Context().Value("db")
				db, ok := temp.(*sql.DB)
				if !ok {
					Srvlog.Error("type assert", "err", ok)
					return
				} else {
					err := DbselectHuser(db, sessionid)
					if err != nil {
						if err == sql.ErrNoRows {
							err := DbinsertHuser(db, sessionid, userPic, nickname)
							if err == nil {
								js_response := simplejson.New()
								js_response.Set("status", 1000)
								js_response.Set("sessionid", sessionid)
								body, err := js_response.Encode()
								if err != nil {
									Srvlog.Error("simplejson MarshalJson", "err", err)
								}
								w.Write(body)
								return
							}
						}
						js_response := simplejson.New()
						js_response.Set("status", 4003)
						js_response.Set("error", "err")
						body, err := js_response.Encode()
						if err != nil {
							Srvlog.Error("simplejson MarshalJson", "err", err.Error())
						} else {
							Srvlog.Info("DbselectHuser", "err", err.Error())
						}
						w.Write(body)
						return
					} else {
						err := DbupdateHuser(db, sessionid, userPic, nickname)
						if err == nil {
							js_response := simplejson.New()
							js_response.Set("status", 1000)
							js_response.Set("sessionid", sessionid)
							body, err := js_response.Encode()
							if err != nil {
								Srvlog.Error("simplejson MarshalJson", "err", err)
							}
							w.Write(body)
							return
						}
					}
				}
			}
		}
	}

}

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(172.17.0.1:3306)/employees")
	if err != nil {
		Srvlog.Crit(err.Error())
		os.Exit(1)
		return
	}
	defer db.Close()
	mux := http.DefaultServeMux
	mux.HandleFunc("/", StatusHandler)
	mux.HandleFunc("/mengyou/login", patter_login(patter_session(Test)))
	mux.HandleFunc("/user/saveOrUpdate", patter_login(patter_session(SaveUserinfo)))
	contextedMux := AddContextSupport(mux, db)
	err = http.ListenAndServeTLS("0.0.0.0:8080", "cert.pem", "key.pem", contextedMux)
	//err = http.ListenAndServe("0.0.0.0:8080", contextedMux)
	Srvlog.Info("check the server status", "err", err)
}
