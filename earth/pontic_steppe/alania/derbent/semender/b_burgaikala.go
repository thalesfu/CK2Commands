package semender

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔格卡拉BurgaikalaBarony struct {
	feud.BaseBarony
}

var BBurgaikala布尔格卡拉 feud.Barony = &布尔格卡拉BurgaikalaBarony{}

func init() {
    f := BBurgaikala布尔格卡拉.(*布尔格卡拉BurgaikalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burgaikala",
		TitleName: "布尔格卡拉",
		TitleCode: "b_burgaikala",
	}
}
