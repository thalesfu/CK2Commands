package lakomelza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车尔科瓦CherkwaBarony struct {
	feud.BaseBarony
}

var BCherkwa车尔科瓦 feud.Barony = &车尔科瓦CherkwaBarony{}

func init() {
    f := BCherkwa车尔科瓦.(*车尔科瓦CherkwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherkwa",
		TitleName: "车尔科瓦",
		TitleCode: "b_cherkwa",
	}
}
