package more

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马德舍MadesjoBarony struct {
	feud.BaseBarony
}

var BMadesjo马德舍 feud.Barony = &马德舍MadesjoBarony{}

func init() {
    f := BMadesjo马德舍.(*马德舍MadesjoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madesjo",
		TitleName: "马德舍",
		TitleCode: "b_madesjo",
	}
}
