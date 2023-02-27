package kaliskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科沃KoloBarony struct {
	feud.BaseBarony
}

var BKolo科沃 feud.Barony = &科沃KoloBarony{}

func init() {
    f := BKolo科沃.(*科沃KoloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolo",
		TitleName: "科沃",
		TitleCode: "b_kolo",
	}
}
