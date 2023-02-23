package pitten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃德尔茨EdelzBarony struct {
	feud.BaseBarony
}

var BEdelz埃德尔茨 feud.Barony = &埃德尔茨EdelzBarony{}

func init() {
	f := BEdelz埃德尔茨.(*埃德尔茨EdelzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "edelz",
		TitleName: "埃德尔茨",
		TitleCode: "b_edelz",
	}
}
