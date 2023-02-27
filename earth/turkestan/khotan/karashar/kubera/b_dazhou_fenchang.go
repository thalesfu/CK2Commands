package kubera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大州分场Dazhou_fenchangBarony struct {
	feud.BaseBarony
}

var BDazhou_fenchang大州分场 feud.Barony = &大州分场Dazhou_fenchangBarony{}

func init() {
    f := BDazhou_fenchang大州分场.(*大州分场Dazhou_fenchangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dazhou_fenchang",
		TitleName: "大州分场",
		TitleCode: "b_dazhou_fenchang",
	}
}
