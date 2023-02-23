package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡奥尔CahorsBarony struct {
	feud.BaseBarony
}

var BCahors卡奥尔 feud.Barony = &卡奥尔CahorsBarony{}

func init() {
	f := BCahors卡奥尔.(*卡奥尔CahorsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cahors",
		TitleName: "卡奥尔",
		TitleCode: "b_cahors",
	}
}
