package data

import (
	"fmt"

	"gopkg.in/jmcvetta/neoism.v1"
)

type NeoDb struct {
	db *neoism.Database
}

func (d NeoDb) Start() {
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		panic(fmt.Sprintf("Could not start db: %s", err))
	}
	d.db = db
}

func (d NeoDb) Create(props neoism.Props) *neoism.Node {
	n, err := d.db.CreateNode(props)
	if err != nil {
		panic(fmt.Sprintf("Could not insert node (%v): %s", props, err))
	}
	return n
}
