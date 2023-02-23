package manyakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘毗罗军荼KovilkondaBarony struct {
	feud.BaseBarony
}

var BKovilkonda拘毗罗军荼 feud.Barony = &拘毗罗军荼KovilkondaBarony{}

func init() {
	f := BKovilkonda拘毗罗军荼.(*拘毗罗军荼KovilkondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kovilkonda",
		TitleName: "拘毗罗军荼",
		TitleCode: "b_kovilkonda",
	}
}
