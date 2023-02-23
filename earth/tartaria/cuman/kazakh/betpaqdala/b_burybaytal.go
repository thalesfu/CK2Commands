package betpaqdala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷尔拜塔尔BurybaytalBarony struct {
	feud.BaseBarony
}

var BBurybaytal布雷尔拜塔尔 feud.Barony = &布雷尔拜塔尔BurybaytalBarony{}

func init() {
	f := BBurybaytal布雷尔拜塔尔.(*布雷尔拜塔尔BurybaytalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burybaytal",
		TitleName: "布雷尔拜塔尔",
		TitleCode: "b_burybaytal",
	}
}
