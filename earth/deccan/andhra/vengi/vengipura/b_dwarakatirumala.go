package vengipura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达瓦拉卡蒂鲁马拉DwarakatirumalaBarony struct {
	feud.BaseBarony
}

var BDwarakatirumala达瓦拉卡蒂鲁马拉 feud.Barony = &达瓦拉卡蒂鲁马拉DwarakatirumalaBarony{}

func init() {
    f := BDwarakatirumala达瓦拉卡蒂鲁马拉.(*达瓦拉卡蒂鲁马拉DwarakatirumalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dwarakatirumala",
		TitleName: "达瓦拉卡蒂鲁马拉",
		TitleCode: "b_dwarakatirumala",
	}
}
