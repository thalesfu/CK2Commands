package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豪特翁HatvanBarony struct {
	feud.BaseBarony
}

var BHatvan豪特翁 feud.Barony = &豪特翁HatvanBarony{}

func init() {
    f := BHatvan豪特翁.(*豪特翁HatvanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hatvan",
		TitleName: "豪特翁",
		TitleCode: "b_hatvan",
	}
}
