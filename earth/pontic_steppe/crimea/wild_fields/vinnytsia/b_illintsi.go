package vinnytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊林齐IllintsiBarony struct {
	feud.BaseBarony
}

var BIllintsi伊林齐 feud.Barony = &伊林齐IllintsiBarony{}

func init() {
    f := BIllintsi伊林齐.(*伊林齐IllintsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "illintsi",
		TitleName: "伊林齐",
		TitleCode: "b_illintsi",
	}
}
