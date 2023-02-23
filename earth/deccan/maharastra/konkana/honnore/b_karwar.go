package honnore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡瓦尔KarwarBarony struct {
	feud.BaseBarony
}

var BKarwar卡瓦尔 feud.Barony = &卡瓦尔KarwarBarony{}

func init() {
	f := BKarwar卡瓦尔.(*卡瓦尔KarwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karwar",
		TitleName: "卡瓦尔",
		TitleCode: "b_karwar",
	}
}
