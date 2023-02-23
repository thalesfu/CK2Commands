package krems

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨韦特尔ZwettlBarony struct {
	feud.BaseBarony
}

var BZwettl茨韦特尔 feud.Barony = &茨韦特尔ZwettlBarony{}

func init() {
	f := BZwettl茨韦特尔.(*茨韦特尔ZwettlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zwettl",
		TitleName: "茨韦特尔",
		TitleCode: "b_zwettl",
	}
}
