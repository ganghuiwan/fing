package main

import (
	"fing"
	"fmt"
	"macdb"
	"strconv"
	"strings"
	"sync"
)

type MB struct {
	ip  string
	mac string
}

var macBuffer = make(chan *MB, 255)
var macCount = make(chan struct{}, 5)
var db *macdb.MacDB

func init() {
	var err error
	db, err = macdb.InitDB("oui.csv")
	if err != nil {
		panic("init db fail " + err.Error())
	}
}

func main() {
	ip, _, err := fing.ExternalIP()
	if err != nil {
		fmt.Println("fing.ExternalIP fail: ", err)
	}
	index := strings.LastIndexByte(ip, '.')
	prefix := ip[:index+1]
	var wg sync.WaitGroup
	for i := 0; i < 256; i++ {
		wg.Add(1)
		macCount <- struct{}{}
		go func(ip string) {
			defer wg.Done()
			defer func() { <-macCount }()
			mac, _, err := fing.Mac(ip)
			if err == nil {
				mb := &MB{ip, mac.String()}
				macBuffer <- mb
			}
		}(prefix + strconv.Itoa(i))
	}
	go func() {
		wg.Wait()
		close(macBuffer)
	}()
	for mac := range macBuffer {
		rsl := db.Get(mac.mac)
		if rsl == nil {
			continue
		}
		fmt.Printf("%15s %17s %s\n", mac.ip, mac.mac, rsl.Organization)
	}
}
