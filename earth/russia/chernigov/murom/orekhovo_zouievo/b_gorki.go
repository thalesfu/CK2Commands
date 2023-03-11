package orekhovo_zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔基GorkiBarony struct {
	feud.BaseBarony
}

var BGorki戈尔基 feud.Barony = &戈尔基GorkiBarony{}

func init() {
    f := BGorki戈尔基.(*戈尔基GorkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorki",
		TitleName: "戈尔基",
		TitleCode: "b_gorki",
	}
}
