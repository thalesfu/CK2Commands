package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛尔卡LorcaBarony struct {
	feud.BaseBarony
}

var BLorca洛尔卡 feud.Barony = &洛尔卡LorcaBarony{}

func init() {
    f := BLorca洛尔卡.(*洛尔卡LorcaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lorca",
		TitleName: "洛尔卡",
		TitleCode: "b_lorca",
	}
}
