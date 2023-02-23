package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔加尔KargalBarony struct {
	feud.BaseBarony
}

var BKargal卡尔加尔 feud.Barony = &卡尔加尔KargalBarony{}

func init() {
	f := BKargal卡尔加尔.(*卡尔加尔KargalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kargal",
		TitleName: "卡尔加尔",
		TitleCode: "b_kargal",
	}
}
