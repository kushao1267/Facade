package utils

import (
	"net/url"
)

func GetComponent(s string) (*url.URL, error){
	u, err := url.Parse(s)
	return u, err
}


func GetHostName(s string) (string, error){
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}