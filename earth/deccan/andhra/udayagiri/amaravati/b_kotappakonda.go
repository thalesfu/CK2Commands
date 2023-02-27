package amaravati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘吒波军荼KotappakondaBarony struct {
	feud.BaseBarony
}

var BKotappakonda拘吒波军荼 feud.Barony = &拘吒波军荼KotappakondaBarony{}

func init() {
    f := BKotappakonda拘吒波军荼.(*拘吒波军荼KotappakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotappakonda",
		TitleName: "拘吒波军荼",
		TitleCode: "b_kotappakonda",
	}
}
