package sunnlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷克雅未克ReykjavikBarony struct {
	feud.BaseBarony
}

var BReykjavik雷克雅未克 feud.Barony = &雷克雅未克ReykjavikBarony{}

func init() {
    f := BReykjavik雷克雅未克.(*雷克雅未克ReykjavikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reykjavik",
		TitleName: "雷克雅未克",
		TitleCode: "b_reykjavik",
	}
}
