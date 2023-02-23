package kunggar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甲玛GyamaBarony struct {
	feud.BaseBarony
}

var BGyama甲玛 feud.Barony = &甲玛GyamaBarony{}

func init() {
	f := BGyama甲玛.(*甲玛GyamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyama",
		TitleName: "甲玛",
		TitleCode: "b_gyama",
	}
}
