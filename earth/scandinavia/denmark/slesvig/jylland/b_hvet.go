package jylland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦特HvetBarony struct {
	feud.BaseBarony
}

var BHvet韦特 feud.Barony = &韦特HvetBarony{}

func init() {
	f := BHvet韦特.(*韦特HvetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hvet",
		TitleName: "韦特",
		TitleCode: "b_hvet",
	}
}
