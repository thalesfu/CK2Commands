package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰雷_卡拉ChallykalaBarony struct {
	feud.BaseBarony
}

var BChallykala恰雷_卡拉 feud.Barony = &恰雷_卡拉ChallykalaBarony{}

func init() {
    f := BChallykala恰雷_卡拉.(*恰雷_卡拉ChallykalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "challykala",
		TitleName: "恰雷_卡拉",
		TitleCode: "b_challykala",
	}
}
