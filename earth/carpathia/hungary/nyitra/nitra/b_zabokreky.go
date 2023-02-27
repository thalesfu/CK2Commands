package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎博克雷基ZabokrekyBarony struct {
	feud.BaseBarony
}

var BZabokreky扎博克雷基 feud.Barony = &扎博克雷基ZabokrekyBarony{}

func init() {
    f := BZabokreky扎博克雷基.(*扎博克雷基ZabokrekyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zabokreky",
		TitleName: "扎博克雷基",
		TitleCode: "b_zabokreky",
	}
}
