package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙雷阿莱MonrealeBarony struct {
	feud.BaseBarony
}

var BMonreale蒙雷阿莱 feud.Barony = &蒙雷阿莱MonrealeBarony{}

func init() {
    f := BMonreale蒙雷阿莱.(*蒙雷阿莱MonrealeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monreale",
		TitleName: "蒙雷阿莱",
		TitleCode: "b_monreale",
	}
}
