package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙日斯卡尔MontgiscardBarony struct {
	feud.BaseBarony
}

var BMontgiscard蒙日斯卡尔 feud.Barony = &蒙日斯卡尔MontgiscardBarony{}

func init() {
	f := BMontgiscard蒙日斯卡尔.(*蒙日斯卡尔MontgiscardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montgiscard",
		TitleName: "蒙日斯卡尔",
		TitleCode: "b_montgiscard",
	}
}
