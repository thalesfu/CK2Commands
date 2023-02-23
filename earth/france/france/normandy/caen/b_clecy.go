package caen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克莱西ClecyBarony struct {
	feud.BaseBarony
}

var BClecy克莱西 feud.Barony = &克莱西ClecyBarony{}

func init() {
	f := BClecy克莱西.(*克莱西ClecyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clecy",
		TitleName: "克莱西",
		TitleCode: "b_clecy",
	}
}
