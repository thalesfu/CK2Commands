package ennedi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨布SabuBarony struct {
	feud.BaseBarony
}

var BSabu萨布 feud.Barony = &萨布SabuBarony{}

func init() {
	f := BSabu萨布.(*萨布SabuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabu",
		TitleName: "萨布",
		TitleCode: "b_sabu",
	}
}
