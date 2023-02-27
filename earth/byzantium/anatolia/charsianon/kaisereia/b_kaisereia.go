package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯撒利亚KaisereiaBarony struct {
	feud.BaseBarony
}

var BKaisereia凯撒利亚 feud.Barony = &凯撒利亚KaisereiaBarony{}

func init() {
    f := BKaisereia凯撒利亚.(*凯撒利亚KaisereiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaisereia",
		TitleName: "凯撒利亚",
		TitleCode: "b_kaisereia",
	}
}
