package main

import "net/url"

func isUrlValid(str string) bool {
	_, err := url.ParseRequestURI(str)
	return nil == err
}

func shouldStop(radios []Radio, forceRefresh bool) bool {
	if len(radios) < pageSize {
		return true
	} else if forceRefresh {
		return false
	}

	for _, radio := range radios {
		data, _ := mgr.Find([]byte(radio.ID))
		if len(data) > 0 {
			return true
		}
	}
	return false
}
