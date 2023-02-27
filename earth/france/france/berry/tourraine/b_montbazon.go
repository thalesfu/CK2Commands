package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙巴宗MontbazonBarony struct {
	feud.BaseBarony
}

var BMontbazon蒙巴宗 feud.Barony = &蒙巴宗MontbazonBarony{}

func init() {
    f := BMontbazon蒙巴宗.(*蒙巴宗MontbazonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montbazon",
		TitleName: "蒙巴宗",
		TitleCode: "b_montbazon",
	}
}
