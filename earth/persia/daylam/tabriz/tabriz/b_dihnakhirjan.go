package tabriz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪纳希尔詹DihnakhirjanBarony struct {
	feud.BaseBarony
}

var BDihnakhirjan迪纳希尔詹 feud.Barony = &迪纳希尔詹DihnakhirjanBarony{}

func init() {
    f := BDihnakhirjan迪纳希尔詹.(*迪纳希尔詹DihnakhirjanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dihnakhirjan",
		TitleName: "迪纳希尔詹",
		TitleCode: "b_dihnakhirjan",
	}
}
