package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姆堡SumburghBarony struct {
	feud.BaseBarony
}

var BSumburgh萨姆堡 feud.Barony = &萨姆堡SumburghBarony{}

func init() {
    f := BSumburgh萨姆堡.(*萨姆堡SumburghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumburgh",
		TitleName: "萨姆堡",
		TitleCode: "b_sumburgh",
	}
}
