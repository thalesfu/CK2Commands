package kursk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法捷日FatejBarony struct {
	feud.BaseBarony
}

var BFatej法捷日 feud.Barony = &法捷日FatejBarony{}

func init() {
    f := BFatej法捷日.(*法捷日FatejBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fatej",
		TitleName: "法捷日",
		TitleCode: "b_fatej",
	}
}
