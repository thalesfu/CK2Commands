package dege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 龚垭GongyaBarony struct {
	feud.BaseBarony
}

var BGongya龚垭 feud.Barony = &龚垭GongyaBarony{}

func init() {
    f := BGongya龚垭.(*龚垭GongyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gongya",
		TitleName: "龚垭",
		TitleCode: "b_gongya",
	}
}
