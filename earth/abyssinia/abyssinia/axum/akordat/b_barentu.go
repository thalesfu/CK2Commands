package akordat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴伦图BarentuBarony struct {
	feud.BaseBarony
}

var BBarentu巴伦图 feud.Barony = &巴伦图BarentuBarony{}

func init() {
    f := BBarentu巴伦图.(*巴伦图BarentuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barentu",
		TitleName: "巴伦图",
		TitleCode: "b_barentu",
	}
}
