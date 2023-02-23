package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙塔内MontanerBarony struct {
	feud.BaseBarony
}

var BMontaner蒙塔内 feud.Barony = &蒙塔内MontanerBarony{}

func init() {
	f := BMontaner蒙塔内.(*蒙塔内MontanerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montaner",
		TitleName: "蒙塔内",
		TitleCode: "b_montaner",
	}
}
