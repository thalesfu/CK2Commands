package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多纳DohnaBarony struct {
	feud.BaseBarony
}

var BDohna多纳 feud.Barony = &多纳DohnaBarony{}

func init() {
    f := BDohna多纳.(*多纳DohnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dohna",
		TitleName: "多纳",
		TitleCode: "b_dohna",
	}
}
