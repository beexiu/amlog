package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type session struct {
	name   string
	addr   string
	expire time.Time
}

type sessions struct {
	reg map[string]*session
	sync.Mutex
}

func ip(r *http.Request) string {
	return strings.SplitN(r.RemoteAddr, ":", 2)[0]
}

func (ss *sessions) Login(w http.ResponseWriter, r *http.Request) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	r.ParseForm()
	user := r.Form.Get("namename")
	pass := r.Form.Get("password")
	if user == "" || pass == "" {
		panic(fmt.Errorf("missing username or password"))
	}

	if !checkUserNameAndPasswd(user, pass) {
		panic(fmt.Errorf("权限验证失败: 用户名或密码错误"))
	}

	sid := uuid(16)
	s := &session{
		name:   user,
		addr:   ip(r),
		expire: time.Now().Add(4 * time.Hour),
	}
	ss.Lock()
	ss.reg[sid] = s
	ss.Unlock()
	setCookie(w, "session", sid, 86400*30)
	setCookie(w, "user", user, 86400*30)
	return
}

func checkUserNameAndPasswd(name string, passwd string) bool {
	if name == "brant" && passwd == "1234" {
		return true
	}

	return false
}

func (ss *sessions) Validate(r *http.Request) bool {
	sid := getCookie(r, "session")
	ss.Lock()
	defer ss.Unlock()
	s := ss.reg[sid]
	if s == nil || s.addr != ip(r) || s.expire.Before(time.Now()) {
		return false
	}
	if s.expire.Sub(time.Now()) < 24*7*time.Hour {
		s.expire = time.Now().Add(24 * 14 * time.Hour)
		ss.reg[sid] = s
	}
	return true
}

func (ss *sessions) CleanUp() {
	ss.Lock()
	defer ss.Unlock()
	for id, s := range ss.reg {
		if time.Now().After(s.expire) {
			delete(ss.reg, id)
		}
	}
}

// SS sessions manager
var SS sessions

func init() {
	SS.reg = make(map[string]*session)
}
