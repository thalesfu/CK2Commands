package shorkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马拉KamaliaBarony struct {
	feud.BaseBarony
}

var BKamalia卡马拉 feud.Barony = &卡马拉KamaliaBarony{}

func init() {
	f := BKamalia卡马拉.(*卡马拉KamaliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamalia",
		TitleName: "卡马拉",
		TitleCode: "b_kamalia",
	}
}
