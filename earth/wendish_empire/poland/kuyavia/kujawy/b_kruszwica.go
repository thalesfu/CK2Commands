package kujawy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁什维察KruszwicaBarony struct {
	feud.BaseBarony
}

var BKruszwica克鲁什维察 feud.Barony = &克鲁什维察KruszwicaBarony{}

func init() {
    f := BKruszwica克鲁什维察.(*克鲁什维察KruszwicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kruszwica",
		TitleName: "克鲁什维察",
		TitleCode: "b_kruszwica",
	}
}
