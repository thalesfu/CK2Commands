package nandapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘罗弗多KoraputBarony struct {
	feud.BaseBarony
}

var BKoraput拘罗弗多 feud.Barony = &拘罗弗多KoraputBarony{}

func init() {
    f := BKoraput拘罗弗多.(*拘罗弗多KoraputBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koraput",
		TitleName: "拘罗弗多",
		TitleCode: "b_koraput",
	}
}
