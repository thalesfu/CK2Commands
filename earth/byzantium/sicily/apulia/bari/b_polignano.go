package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利尼亚诺PolignanoBarony struct {
	feud.BaseBarony
}

var BPolignano波利尼亚诺 feud.Barony = &波利尼亚诺PolignanoBarony{}

func init() {
	f := BPolignano波利尼亚诺.(*波利尼亚诺PolignanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polignano",
		TitleName: "波利尼亚诺",
		TitleCode: "b_polignano",
	}
}
