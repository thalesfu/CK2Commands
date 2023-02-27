package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_阿克赛斯卡亚UstaksayaskayaBarony struct {
	feud.BaseBarony
}

var BUstaksayaskaya乌斯季_阿克赛斯卡亚 feud.Barony = &乌斯季_阿克赛斯卡亚UstaksayaskayaBarony{}

func init() {
    f := BUstaksayaskaya乌斯季_阿克赛斯卡亚.(*乌斯季_阿克赛斯卡亚UstaksayaskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustaksayaskaya",
		TitleName: "乌斯季_阿克赛斯卡亚",
		TitleCode: "b_ustaksayaskaya",
	}
}
