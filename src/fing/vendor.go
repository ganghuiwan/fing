package fing

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/franela/goreq"
)

func Vendor(mac string) (string, error) {
	macs := strings.Split(mac, ":")
	if len(macs) != 6 {
		return "", fmt.Errorf("MAC Error: %s", mac)
	}
	mac = strings.Join(macs[0:3], "-")
	form := url.Values{}
	form.Add("x", mac)
	form.Add("submit2", "Search!")
	fmt.Println(form.Encode())
	res, err := goreq.Request{
		Method:      "POST",
		Uri:         "http://standards.ieee.org/cgi-bin/ouisearch",
		ContentType: "application/x-www-form-urlencoded",
		UserAgent:   "Cyeam",
		ShowDebug:   false,
		Body:        form.Encode(),
	}.Do()
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := res.Body.ToString()
	if err != nil || body == "" {
		return "", err
	}
	fmt.Println("body=", body)
	vendor := body[strings.Index(body, strings.ToUpper(mac))+len(mac):]
	vendor = strings.TrimLeft(vendor, "</b>   (hex)")
	vendor = strings.TrimSpace(vendor)
	return strings.Split(vendor, "\n")[0], nil
}
