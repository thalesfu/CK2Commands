package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨摩斯SamosBarony struct {
	feud.BaseBarony
}

var BSamos萨摩斯 feud.Barony = &萨摩斯SamosBarony{}

func init() {
    f := BSamos萨摩斯.(*萨摩斯SamosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samos",
		TitleName: "萨摩斯",
		TitleCode: "b_samos",
	}
}
