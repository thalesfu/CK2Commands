package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德里安库尔DrincourtBarony struct {
	feud.BaseBarony
}

var BDrincourt德里安库尔 feud.Barony = &德里安库尔DrincourtBarony{}

func init() {
	f := BDrincourt德里安库尔.(*德里安库尔DrincourtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drincourt",
		TitleName: "德里安库尔",
		TitleCode: "b_drincourt",
	}
}
