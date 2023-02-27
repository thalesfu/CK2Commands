package gurganj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉捷佩KaradepeBarony struct {
	feud.BaseBarony
}

var BKaradepe卡拉捷佩 feud.Barony = &卡拉捷佩KaradepeBarony{}

func init() {
    f := BKaradepe卡拉捷佩.(*卡拉捷佩KaradepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karadepe",
		TitleName: "卡拉捷佩",
		TitleCode: "b_karadepe",
	}
}
