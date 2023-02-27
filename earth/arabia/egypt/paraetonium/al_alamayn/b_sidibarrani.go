package al_alamayn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪拜拉尼SidibarraniBarony struct {
	feud.BaseBarony
}

var BSidibarrani西迪拜拉尼 feud.Barony = &西迪拜拉尼SidibarraniBarony{}

func init() {
    f := BSidibarrani西迪拜拉尼.(*西迪拜拉尼SidibarraniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidibarrani",
		TitleName: "西迪拜拉尼",
		TitleCode: "b_sidibarrani",
	}
}
