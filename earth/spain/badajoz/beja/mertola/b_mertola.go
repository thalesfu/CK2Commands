package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔图拉MertolaBarony struct {
	feud.BaseBarony
}

var BMertola梅尔图拉 feud.Barony = &梅尔图拉MertolaBarony{}

func init() {
    f := BMertola梅尔图拉.(*梅尔图拉MertolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mertola",
		TitleName: "梅尔图拉",
		TitleCode: "b_mertola",
	}
}
