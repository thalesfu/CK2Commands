package kinda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希萨KisaBarony struct {
	feud.BaseBarony
}

var BKisa希萨 feud.Barony = &希萨KisaBarony{}

func init() {
    f := BKisa希萨.(*希萨KisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kisa",
		TitleName: "希萨",
		TitleCode: "b_kisa",
	}
}
