package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔奇克DobrichBarony struct {
	feud.BaseBarony
}

var BDobrich巴尔奇克 feud.Barony = &巴尔奇克DobrichBarony{}

func init() {
	f := BDobrich巴尔奇克.(*巴尔奇克DobrichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobrich",
		TitleName: "巴尔奇克",
		TitleCode: "b_dobrich",
	}
}
