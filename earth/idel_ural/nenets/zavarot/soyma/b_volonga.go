package soyma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃隆加VolongaBarony struct {
	feud.BaseBarony
}

var BVolonga沃隆加 feud.Barony = &沃隆加VolongaBarony{}

func init() {
    f := BVolonga沃隆加.(*沃隆加VolongaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volonga",
		TitleName: "沃隆加",
		TitleCode: "b_volonga",
	}
}
