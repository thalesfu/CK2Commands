package kagha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科托科KotokoBarony struct {
	feud.BaseBarony
}

var BKotoko科托科 feud.Barony = &科托科KotokoBarony{}

func init() {
    f := BKotoko科托科.(*科托科KotokoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotoko",
		TitleName: "科托科",
		TitleCode: "b_kotoko",
	}
}
