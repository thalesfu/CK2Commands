package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫拉MouraBarony struct {
	feud.BaseBarony
}

var BMoura莫拉 feud.Barony = &莫拉MouraBarony{}

func init() {
	f := BMoura莫拉.(*莫拉MouraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moura",
		TitleName: "莫拉",
		TitleCode: "b_moura",
	}
}
