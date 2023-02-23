package chuy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 裴罗将军城BalasagunBarony struct {
	feud.BaseBarony
}

var BBalasagun裴罗将军城 feud.Barony = &裴罗将军城BalasagunBarony{}

func init() {
	f := BBalasagun裴罗将军城.(*裴罗将军城BalasagunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balasagun",
		TitleName: "裴罗将军城",
		TitleCode: "b_balasagun",
	}
}
