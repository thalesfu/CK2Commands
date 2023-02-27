package kozhva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科日瓦KozhvaBarony struct {
	feud.BaseBarony
}

var BKozhva科日瓦 feud.Barony = &科日瓦KozhvaBarony{}

func init() {
    f := BKozhva科日瓦.(*科日瓦KozhvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kozhva",
		TitleName: "科日瓦",
		TitleCode: "b_kozhva",
	}
}
