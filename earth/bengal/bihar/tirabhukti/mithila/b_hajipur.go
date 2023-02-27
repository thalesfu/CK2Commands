package mithila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃耆城HajipurBarony struct {
	feud.BaseBarony
}

var BHajipur诃耆城 feud.Barony = &诃耆城HajipurBarony{}

func init() {
    f := BHajipur诃耆城.(*诃耆城HajipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hajipur",
		TitleName: "诃耆城",
		TitleCode: "b_hajipur",
	}
}
