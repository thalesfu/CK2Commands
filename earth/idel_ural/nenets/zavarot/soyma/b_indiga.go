package soyma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因迪加IndigaBarony struct {
	feud.BaseBarony
}

var BIndiga因迪加 feud.Barony = &因迪加IndigaBarony{}

func init() {
    f := BIndiga因迪加.(*因迪加IndigaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "indiga",
		TitleName: "因迪加",
		TitleCode: "b_indiga",
	}
}
