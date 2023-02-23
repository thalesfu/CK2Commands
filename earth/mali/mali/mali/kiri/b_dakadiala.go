package kiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达卡迪亚拉DakadialaBarony struct {
	feud.BaseBarony
}

var BDakadiala达卡迪亚拉 feud.Barony = &达卡迪亚拉DakadialaBarony{}

func init() {
	f := BDakadiala达卡迪亚拉.(*达卡迪亚拉DakadialaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dakadiala",
		TitleName: "达卡迪亚拉",
		TitleCode: "b_dakadiala",
	}
}
