package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马德里MadridBarony struct {
	feud.BaseBarony
}

var BMadrid马德里 feud.Barony = &马德里MadridBarony{}

func init() {
	f := BMadrid马德里.(*马德里MadridBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madrid",
		TitleName: "马德里",
		TitleCode: "b_madrid",
	}
}
