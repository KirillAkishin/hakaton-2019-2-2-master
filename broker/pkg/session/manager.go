package session

import (
	"broker/pkg/clients"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var (
	tokenSecret = []byte("hakaton-2019-2-2")
)

type SessionsManager struct {
	clientRepo clients.ClientRepo
}

func NewSessionsManager(clientRepo clients.ClientRepo) *SessionsManager {
	return &SessionsManager{
		clientRepo: clientRepo,
	}
}

func (sm *SessionsManager) Check(r *http.Request) (*clients.Client, error) {
	sessionCookie, err := r.Cookie("token")
	if err == http.ErrNoCookie || sessionCookie == nil {
		return nil, ErrNoAuth
	}
	inToken := sessionCookie.Value

	hashSecretGetter := func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			return nil, fmt.Errorf("bad sign method")
		}
		return tokenSecret, nil
	}
	token, err := jwt.Parse(inToken, hashSecretGetter)
	if err != nil || !token.Valid {
		return nil, ErrNoAuth
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrNoAuth
	}
	clientRaw, found := payload["client"]
	if !found {
		return nil, ErrNoAuth
	}
	clientMap, ok := clientRaw.(clients.Client)
	if !ok {
		return nil, ErrNoAuth
	}
	clientID := clientMap.ID
	client, err := sm.clientRepo.GetByID(clientID)
	if err != nil {
		return nil, ErrNoAuth
	}
	return client, nil
}

func (sm *SessionsManager) Create(w http.ResponseWriter, r *http.Request, client *clients.Client) (string, error) {
	iat := time.Now()
	exp := iat.Add(time.Hour * 24 * 7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client": &clients.Client{
			Name:     client.Name,
			ID:       client.ID,
			Balance:  client.Balance,
		},
		"iat": iat.Unix(),
		"exp": exp.Unix(),
	})
	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err
	}
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: exp,
		Path:    "/",
	}
	http.SetCookie(w, cookie)
	return tokenString, nil
}
