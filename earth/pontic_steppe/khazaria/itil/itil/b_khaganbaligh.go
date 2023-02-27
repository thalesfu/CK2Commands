package itil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 可汗八里KhaganbalighBarony struct {
	feud.BaseBarony
}

var BKhaganbaligh可汗八里 feud.Barony = &可汗八里KhaganbalighBarony{}

func init() {
    f := BKhaganbaligh可汗八里.(*可汗八里KhaganbalighBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khaganbaligh",
		TitleName: "可汗八里",
		TitleCode: "b_khaganbaligh",
	}
}
