package emba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴卡奇BakachiBarony struct {
	feud.BaseBarony
}

var BBakachi巴卡奇 feud.Barony = &巴卡奇BakachiBarony{}

func init() {
	f := BBakachi巴卡奇.(*巴卡奇BakachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakachi",
		TitleName: "巴卡奇",
		TitleCode: "b_bakachi",
	}
}
