package handlers

import (
	"broker/pkg/clients"
	orderHistory "broker/pkg/order_history"
	"broker/pkg/requests"
	"broker/pkg/session"
	"broker/pkg/stats"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type ClientHandler struct {
	Tmpl             *template.Template
	ClientRepo       clients.ClientRepo
	OrderHistoryRepo orderHistory.OrderHistoryRepo
	StatsRepo        stats.StatsRepo
	ReqestsRepo      requests.RequestsRepo
	Sessions         *session.SessionsManager
}

func (handler *ClientHandler) Template(w http.ResponseWriter, r *http.Request) {
	handler.Tmpl.ExecuteTemplate(w, "index.html", struct{}{})
}

func (handler *ClientHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	clientForm := clients.ClientForm{}
	err := json.Unmarshal(body, &clientForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Client, Err := handler.ClientRepo.GetByName(clientForm.Name)
	if Err == clients.ErrNoUser {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if Client.Password != clientForm.Pass {
		w.WriteHeader(http.StatusLocked)
		return
	}
	_, err = handler.Sessions.Create(w, r, Client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *ClientHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	clientForm := clients.ClientForm{}
	err := json.Unmarshal(body, &clientForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = handler.ClientRepo.GetByName(clientForm.Name)
	if err == nil {
		w.WriteHeader(http.StatusLocked)
		return
	}
	handler.ClientRepo.Register(&clientForm)

}

func (handler *ClientHandler) Deal(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	dealForm := &requests.DealForm{}
	err := json.Unmarshal(body, dealForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client, err := session.ClientFromContext(r.Context())
	err = handler.ReqestsRepo.Add(dealForm, client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *ClientHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	idForm := IdForm{}
	err := json.Unmarshal(body, &idForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *ClientHandler) History(w http.ResponseWriter, r *http.Request) {
	ticker := r.FormValue("ticker")
	allStat, err := handler.StatsRepo.GetAllByTicker(ticker)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	byteStat, err := json.Marshal(allStat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(byteStat)
}

func (handler *ClientHandler) Stats(w http.ResponseWriter, r *http.Request) {
	ordHist, err := handler.OrderHistoryRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	byteHist, err := json.Marshal(ordHist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(byteHist)

}
