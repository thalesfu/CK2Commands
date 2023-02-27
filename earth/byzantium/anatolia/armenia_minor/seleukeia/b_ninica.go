package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼尼察NinicaBarony struct {
	feud.BaseBarony
}

var BNinica尼尼察 feud.Barony = &尼尼察NinicaBarony{}

func init() {
    f := BNinica尼尼察.(*尼尼察NinicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ninica",
		TitleName: "尼尼察",
		TitleCode: "b_ninica",
	}
}
