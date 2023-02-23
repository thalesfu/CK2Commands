package otuken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于都斤OtukenBarony struct {
	feud.BaseBarony
}

var BOtuken于都斤 feud.Barony = &于都斤OtukenBarony{}

func init() {
	f := BOtuken于都斤.(*于都斤OtukenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otuken",
		TitleName: "于都斤",
		TitleCode: "b_otuken",
	}
}
