package lhunze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日当RitangBarony struct {
	feud.BaseBarony
}

var BRitang日当 feud.Barony = &日当RitangBarony{}

func init() {
    f := BRitang日当.(*日当RitangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ritang",
		TitleName: "日当",
		TitleCode: "b_ritang",
	}
}
