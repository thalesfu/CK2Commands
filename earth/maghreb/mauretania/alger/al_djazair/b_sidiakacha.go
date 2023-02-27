package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪阿卡沙SidiakachaBarony struct {
	feud.BaseBarony
}

var BSidiakacha西迪阿卡沙 feud.Barony = &西迪阿卡沙SidiakachaBarony{}

func init() {
    f := BSidiakacha西迪阿卡沙.(*西迪阿卡沙SidiakachaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidiakacha",
		TitleName: "西迪阿卡沙",
		TitleCode: "b_sidiakacha",
	}
}
