package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕塔拉PataraBarony struct {
	feud.BaseBarony
}

var BPatara帕塔拉 feud.Barony = &帕塔拉PataraBarony{}

func init() {
    f := BPatara帕塔拉.(*帕塔拉PataraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patara",
		TitleName: "帕塔拉",
		TitleCode: "b_patara",
	}
}
