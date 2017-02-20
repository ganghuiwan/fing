package macdb

import "strings"

type Mac struct {
	Registry     string
	Assignment   string
	Organization string
	Address      string
}

type MacDB struct {
	db map[string]*Mac
}

func (mdb *MacDB) Add(mac *Mac) {
	mdb.db[strings.ToLower(mac.Assignment)] = mac
}

func (mdb *MacDB) Get(mac string) *Mac {
	mac = strings.Replace(mac[:8], ":", "", -1)
	return mdb.db[mac]
}

func (mdb *MacDB) Len() int {
	return len(mdb.db)
}
