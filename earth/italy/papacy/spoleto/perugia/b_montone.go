package perugia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙托内MontoneBarony struct {
	feud.BaseBarony
}

var BMontone蒙托内 feud.Barony = &蒙托内MontoneBarony{}

func init() {
	f := BMontone蒙托内.(*蒙托内MontoneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montone",
		TitleName: "蒙托内",
		TitleCode: "b_montone",
	}
}
