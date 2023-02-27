package kundina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼旬罗NijharBarony struct {
	feud.BaseBarony
}

var BNijhar尼旬罗 feud.Barony = &尼旬罗NijharBarony{}

func init() {
    f := BNijhar尼旬罗.(*尼旬罗NijharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nijhar",
		TitleName: "尼旬罗",
		TitleCode: "b_nijhar",
	}
}
