package katun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔巴扎TurbazaBarony struct {
	feud.BaseBarony
}

var BTurbaza图尔巴扎 feud.Barony = &图尔巴扎TurbazaBarony{}

func init() {
	f := BTurbaza图尔巴扎.(*图尔巴扎TurbazaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turbaza",
		TitleName: "图尔巴扎",
		TitleCode: "b_turbaza",
	}
}
