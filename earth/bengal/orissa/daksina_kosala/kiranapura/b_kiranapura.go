package kiranapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日光城KiranapuraBarony struct {
	feud.BaseBarony
}

var BKiranapura日光城 feud.Barony = &日光城KiranapuraBarony{}

func init() {
    f := BKiranapura日光城.(*日光城KiranapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiranapura",
		TitleName: "日光城",
		TitleCode: "b_kiranapura",
	}
}
