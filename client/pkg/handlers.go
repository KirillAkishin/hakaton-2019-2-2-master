package pkg

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type UserHandler struct {
	Tmpl     *template.Template
	UserName string
	UserRepo *user.UserRepo
	Sessions *session.SessionsManager
}

func (handler *UserHandler) Template(w http.ResponseWriter, r *http.Request) {
	handler.Tmpl.ExecuteTemplate(w, "index.html", struct{}{})
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	preUser := PreUser{}
	err := json.Unmarshal(body, &preUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exist, curUser := handler.UserRepo.UserExist(preUser.UserName)
	if !exist {
		mes := system.Message{Message: "user not found"}
		jsonMes, err := mes.ToJson()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonMes)
		return
	}
	if !curUser.CheckPassword(preUser.Password) {
		mes := system.Message{Message: "invalid password"}
		jsonMes, err := mes.ToJson()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonMes)
		return
	}
	jsonToken, err := curUser.JsonToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Sessions.CreateSession(w, curUser.ID, curUser.UserName)
	w.Write(jsonToken)
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	preUser := PreUser{}
	err := json.Unmarshal(body, &preUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exist, curUser := handler.UserRepo.UserExist(preUser.UserName)
	if exist {
		newError := system.Error{
			Location: "body",
			Param:    "username",
			Value:    curUser.UserName,
			Message:  "already exists",
		}
		errors := system.Errors{Errors: []system.Error{newError}}
		jsonMes, err := errors.ToJson()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(422)
		w.Write(jsonMes)
		return
	}
	newUser := user.NewUser(preUser.UserName, preUser.Password)
	jsonToken, err := newUser.JsonToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.UserRepo.AddUser(&newUser)
	handler.Sessions.CreateSession(w, newUser.ID, newUser.UserName)
	w.Write(jsonToken)
}
