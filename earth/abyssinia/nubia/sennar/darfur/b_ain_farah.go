package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因法拉赫Ain_farahBarony struct {
	feud.BaseBarony
}

var BAin_farah艾因法拉赫 feud.Barony = &艾因法拉赫Ain_farahBarony{}

func init() {
    f := BAin_farah艾因法拉赫.(*艾因法拉赫Ain_farahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ain_farah",
		TitleName: "艾因法拉赫",
		TitleCode: "b_ain_farah",
	}
}
