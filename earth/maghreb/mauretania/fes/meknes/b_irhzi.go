package meknes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔齐IrhziBarony struct {
	feud.BaseBarony
}

var BIrhzi伊尔齐 feud.Barony = &伊尔齐IrhziBarony{}

func init() {
    f := BIrhzi伊尔齐.(*伊尔齐IrhziBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irhzi",
		TitleName: "伊尔齐",
		TitleCode: "b_irhzi",
	}
}
