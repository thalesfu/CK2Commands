package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙塔日MontargisBarony struct {
	feud.BaseBarony
}

var BMontargis蒙塔日 feud.Barony = &蒙塔日MontargisBarony{}

func init() {
	f := BMontargis蒙塔日.(*蒙塔日MontargisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montargis",
		TitleName: "蒙塔日",
		TitleCode: "b_montargis",
	}
}
