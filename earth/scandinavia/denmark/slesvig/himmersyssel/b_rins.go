package himmersyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林斯RinsBarony struct {
	feud.BaseBarony
}

var BRins林斯 feud.Barony = &林斯RinsBarony{}

func init() {
	f := BRins林斯.(*林斯RinsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rins",
		TitleName: "林斯",
		TitleCode: "b_rins",
	}
}
