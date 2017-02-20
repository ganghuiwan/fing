package macdb

import (
	"encoding/csv"
	"io"
	"os"
)

func InitDB(fileName string) (*MacDB, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	read := csv.NewReader(file)
	db := &MacDB{make(map[string]*Mac)}
	for {
		record, err := read.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if len(record) != 4 {
			continue
		}
		db.Add(&Mac{record[0], record[1], record[2], record[3]})
	}
	return db, nil
}
