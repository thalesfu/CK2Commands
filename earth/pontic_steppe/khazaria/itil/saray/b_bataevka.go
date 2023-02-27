package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴塔耶夫卡BataevkaBarony struct {
	feud.BaseBarony
}

var BBataevka巴塔耶夫卡 feud.Barony = &巴塔耶夫卡BataevkaBarony{}

func init() {
    f := BBataevka巴塔耶夫卡.(*巴塔耶夫卡BataevkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bataevka",
		TitleName: "巴塔耶夫卡",
		TitleCode: "b_bataevka",
	}
}
