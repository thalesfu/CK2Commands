package bambuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪乌博耶DiouboyeBarony struct {
	feud.BaseBarony
}

var BDiouboye迪乌博耶 feud.Barony = &迪乌博耶DiouboyeBarony{}

func init() {
    f := BDiouboye迪乌博耶.(*迪乌博耶DiouboyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diouboye",
		TitleName: "迪乌博耶",
		TitleCode: "b_diouboye",
	}
}
