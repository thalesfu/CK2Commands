package naldurg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔马拉KarmalaBarony struct {
	feud.BaseBarony
}

var BKarmala卡尔马拉 feud.Barony = &卡尔马拉KarmalaBarony{}

func init() {
	f := BKarmala卡尔马拉.(*卡尔马拉KarmalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karmala",
		TitleName: "卡尔马拉",
		TitleCode: "b_karmala",
	}
}
