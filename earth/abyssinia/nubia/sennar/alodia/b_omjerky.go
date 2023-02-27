package alodia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧米茄尔可OmjerkyBarony struct {
	feud.BaseBarony
}

var BOmjerky欧米茄尔可 feud.Barony = &欧米茄尔可OmjerkyBarony{}

func init() {
    f := BOmjerky欧米茄尔可.(*欧米茄尔可OmjerkyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "omjerky",
		TitleName: "欧米茄尔可",
		TitleCode: "b_omjerky",
	}
}
