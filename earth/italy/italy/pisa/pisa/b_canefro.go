package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡内夫罗CanefroBarony struct {
	feud.BaseBarony
}

var BCanefro卡内夫罗 feud.Barony = &卡内夫罗CanefroBarony{}

func init() {
    f := BCanefro卡内夫罗.(*卡内夫罗CanefroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "canefro",
		TitleName: "卡内夫罗",
		TitleCode: "b_canefro",
	}
}
