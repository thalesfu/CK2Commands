package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼赫底NehardeaBarony struct {
	feud.BaseBarony
}

var BNehardea尼赫底 feud.Barony = &尼赫底NehardeaBarony{}

func init() {
	f := BNehardea尼赫底.(*尼赫底NehardeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nehardea",
		TitleName: "尼赫底",
		TitleCode: "b_nehardea",
	}
}
