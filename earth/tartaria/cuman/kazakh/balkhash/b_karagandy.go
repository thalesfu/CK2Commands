package balkhash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉干达KaragandyBarony struct {
	feud.BaseBarony
}

var BKaragandy卡拉干达 feud.Barony = &卡拉干达KaragandyBarony{}

func init() {
	f := BKaragandy卡拉干达.(*卡拉干达KaragandyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karagandy",
		TitleName: "卡拉干达",
		TitleCode: "b_karagandy",
	}
}
