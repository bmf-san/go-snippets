package main

import (
	"encoding/base64"
	"net/http"
	"time"
)

// SetFlash set a flash message.
func SetFlash(w http.ResponseWriter, name string, value []byte) {
	http.SetCookie(w, &http.Cookie{
		Name:  name,
		Value: base64.RawURLEncoding.EncodeToString(value),
	})
}

// GetFlash get a flash message from cookie.
func GetFlash(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}
	value, err := base64.RawURLEncoding.DecodeString(c.Value)
	if err != nil {
		return nil, err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	})
	return value, nil
}
