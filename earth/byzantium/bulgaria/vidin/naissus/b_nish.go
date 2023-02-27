package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼什NishBarony struct {
	feud.BaseBarony
}

var BNish尼什 feud.Barony = &尼什NishBarony{}

func init() {
    f := BNish尼什.(*尼什NishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nish",
		TitleName: "尼什",
		TitleCode: "b_nish",
	}
}
