package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别洛戈尔斯克BelogorskBarony struct {
	feud.BaseBarony
}

var BBelogorsk别洛戈尔斯克 feud.Barony = &别洛戈尔斯克BelogorskBarony{}

func init() {
	f := BBelogorsk别洛戈尔斯克.(*别洛戈尔斯克BelogorskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belogorsk",
		TitleName: "别洛戈尔斯克",
		TitleCode: "b_belogorsk",
	}
}
