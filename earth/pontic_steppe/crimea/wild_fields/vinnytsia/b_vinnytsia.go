package vinnytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 文尼察VinnytsiaBarony struct {
	feud.BaseBarony
}

var BVinnytsia文尼察 feud.Barony = &文尼察VinnytsiaBarony{}

func init() {
    f := BVinnytsia文尼察.(*文尼察VinnytsiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vinnytsia",
		TitleName: "文尼察",
		TitleCode: "b_vinnytsia",
	}
}
