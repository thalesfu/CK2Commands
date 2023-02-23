package bilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达努瓦拉DanoualaBarony struct {
	feud.BaseBarony
}

var BDanouala达努瓦拉 feud.Barony = &达努瓦拉DanoualaBarony{}

func init() {
	f := BDanouala达努瓦拉.(*达努瓦拉DanoualaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danouala",
		TitleName: "达努瓦拉",
		TitleCode: "b_danouala",
	}
}
