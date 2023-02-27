package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙蒂雷亚尔MonterealBarony struct {
	feud.BaseBarony
}

var BMontereal蒙蒂雷亚尔 feud.Barony = &蒙蒂雷亚尔MonterealBarony{}

func init() {
    f := BMontereal蒙蒂雷亚尔.(*蒙蒂雷亚尔MonterealBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montereal",
		TitleName: "蒙蒂雷亚尔",
		TitleCode: "b_montereal",
	}
}
