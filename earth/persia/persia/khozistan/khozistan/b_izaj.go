package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊泽IzajBarony struct {
	feud.BaseBarony
}

var BIzaj伊泽 feud.Barony = &伊泽IzajBarony{}

func init() {
    f := BIzaj伊泽.(*伊泽IzajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "izaj",
		TitleName: "伊泽",
		TitleCode: "b_izaj",
	}
}
