package regensburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷根斯堡RegensburgBarony struct {
	feud.BaseBarony
}

var BRegensburg雷根斯堡 feud.Barony = &雷根斯堡RegensburgBarony{}

func init() {
    f := BRegensburg雷根斯堡.(*雷根斯堡RegensburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "regensburg",
		TitleName: "雷根斯堡",
		TitleCode: "b_regensburg",
	}
}
