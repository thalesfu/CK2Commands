package kaliskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科宁KoninBarony struct {
	feud.BaseBarony
}

var BKonin科宁 feud.Barony = &科宁KoninBarony{}

func init() {
    f := BKonin科宁.(*科宁KoninBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konin",
		TitleName: "科宁",
		TitleCode: "b_konin",
	}
}
