package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛_厄捷斯SayutesBarony struct {
	feud.BaseBarony
}

var BSayutes赛_厄捷斯 feud.Barony = &赛_厄捷斯SayutesBarony{}

func init() {
	f := BSayutes赛_厄捷斯.(*赛_厄捷斯SayutesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sayutes",
		TitleName: "赛_厄捷斯",
		TitleCode: "b_sayutes",
	}
}
