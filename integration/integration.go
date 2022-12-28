package integration

import (
	"net/http"
)

func TestaSite(site string) (int, error) {
	res, err := http.Get(site)
	return res.StatusCode, err
}
