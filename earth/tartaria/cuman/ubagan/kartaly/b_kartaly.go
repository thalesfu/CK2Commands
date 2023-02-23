package kartaly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔塔雷_阿亚特KartalyBarony struct {
	feud.BaseBarony
}

var BKartaly卡尔塔雷_阿亚特 feud.Barony = &卡尔塔雷_阿亚特KartalyBarony{}

func init() {
	f := BKartaly卡尔塔雷_阿亚特.(*卡尔塔雷_阿亚特KartalyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kartaly",
		TitleName: "卡尔塔雷_阿亚特",
		TitleCode: "b_kartaly",
	}
}
