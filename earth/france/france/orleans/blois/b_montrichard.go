package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特里夏尔MontrichardBarony struct {
	feud.BaseBarony
}

var BMontrichard蒙特里夏尔 feud.Barony = &蒙特里夏尔MontrichardBarony{}

func init() {
	f := BMontrichard蒙特里夏尔.(*蒙特里夏尔MontrichardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montrichard",
		TitleName: "蒙特里夏尔",
		TitleCode: "b_montrichard",
	}
}
