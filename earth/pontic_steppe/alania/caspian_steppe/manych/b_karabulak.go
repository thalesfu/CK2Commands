package manych

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉布拉克KarabulakBarony struct {
	feud.BaseBarony
}

var BKarabulak卡拉布拉克 feud.Barony = &卡拉布拉克KarabulakBarony{}

func init() {
    f := BKarabulak卡拉布拉克.(*卡拉布拉克KarabulakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karabulak",
		TitleName: "卡拉布拉克",
		TitleCode: "b_karabulak",
	}
}
