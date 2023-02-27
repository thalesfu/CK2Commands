package mtskheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德瓦力DjvariBarony struct {
	feud.BaseBarony
}

var BDjvari德瓦力 feud.Barony = &德瓦力DjvariBarony{}

func init() {
    f := BDjvari德瓦力.(*德瓦力DjvariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djvari",
		TitleName: "德瓦力",
		TitleCode: "b_djvari",
	}
}
