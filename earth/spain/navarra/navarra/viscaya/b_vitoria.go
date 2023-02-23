package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维多利亚VitoriaBarony struct {
	feud.BaseBarony
}

var BVitoria维多利亚 feud.Barony = &维多利亚VitoriaBarony{}

func init() {
	f := BVitoria维多利亚.(*维多利亚VitoriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vitoria",
		TitleName: "维多利亚",
		TitleCode: "b_vitoria",
	}
}
