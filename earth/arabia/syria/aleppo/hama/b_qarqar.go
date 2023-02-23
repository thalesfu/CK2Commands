package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甲加QarqarBarony struct {
	feud.BaseBarony
}

var BQarqar甲加 feud.Barony = &甲加QarqarBarony{}

func init() {
	f := BQarqar甲加.(*甲加QarqarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qarqar",
		TitleName: "甲加",
		TitleCode: "b_qarqar",
	}
}
