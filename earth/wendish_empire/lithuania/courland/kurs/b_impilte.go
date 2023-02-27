package kurs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊皮尔蒂斯ImpilteBarony struct {
	feud.BaseBarony
}

var BImpilte伊皮尔蒂斯 feud.Barony = &伊皮尔蒂斯ImpilteBarony{}

func init() {
    f := BImpilte伊皮尔蒂斯.(*伊皮尔蒂斯ImpilteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "impilte",
		TitleName: "伊皮尔蒂斯",
		TitleCode: "b_impilte",
	}
}
