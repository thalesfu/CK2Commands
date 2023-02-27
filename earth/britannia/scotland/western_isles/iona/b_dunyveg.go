package iona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丹尼维哥堡DunyvegBarony struct {
	feud.BaseBarony
}

var BDunyveg丹尼维哥堡 feud.Barony = &丹尼维哥堡DunyvegBarony{}

func init() {
    f := BDunyveg丹尼维哥堡.(*丹尼维哥堡DunyvegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunyveg",
		TitleName: "丹尼维哥堡",
		TitleCode: "b_dunyveg",
	}
}
