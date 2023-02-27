package lhatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡嘎KagarBarony struct {
	feud.BaseBarony
}

var BKagar卡嘎 feud.Barony = &卡嘎KagarBarony{}

func init() {
    f := BKagar卡嘎.(*卡嘎KagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kagar",
		TitleName: "卡嘎",
		TitleCode: "b_kagar",
	}
}
