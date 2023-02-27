package kujawy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克沃达瓦KlodawaBarony struct {
	feud.BaseBarony
}

var BKlodawa克沃达瓦 feud.Barony = &克沃达瓦KlodawaBarony{}

func init() {
    f := BKlodawa克沃达瓦.(*克沃达瓦KlodawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klodawa",
		TitleName: "克沃达瓦",
		TitleCode: "b_klodawa",
	}
}
