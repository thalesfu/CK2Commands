package rutog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日托RutogBarony struct {
	feud.BaseBarony
}

var BRutog日托 feud.Barony = &日托RutogBarony{}

func init() {
	f := BRutog日托.(*日托RutogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rutog",
		TitleName: "日托",
		TitleCode: "b_rutog",
	}
}
