package kolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔瓦KolvaBarony struct {
	feud.BaseBarony
}

var BKolva科尔瓦 feud.Barony = &科尔瓦KolvaBarony{}

func init() {
    f := BKolva科尔瓦.(*科尔瓦KolvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolva",
		TitleName: "科尔瓦",
		TitleCode: "b_kolva",
	}
}
