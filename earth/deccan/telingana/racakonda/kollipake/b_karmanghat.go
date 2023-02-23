package kollipake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔曼哈特KarmanghatBarony struct {
	feud.BaseBarony
}

var BKarmanghat卡尔曼哈特 feud.Barony = &卡尔曼哈特KarmanghatBarony{}

func init() {
	f := BKarmanghat卡尔曼哈特.(*卡尔曼哈特KarmanghatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karmanghat",
		TitleName: "卡尔曼哈特",
		TitleCode: "b_karmanghat",
	}
}
