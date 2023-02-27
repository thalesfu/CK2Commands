package sthanisvara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尼跋PanipatBarony struct {
	feud.BaseBarony
}

var BPanipat波尼跋 feud.Barony = &波尼跋PanipatBarony{}

func init() {
    f := BPanipat波尼跋.(*波尼跋PanipatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panipat",
		TitleName: "波尼跋",
		TitleCode: "b_panipat",
	}
}
