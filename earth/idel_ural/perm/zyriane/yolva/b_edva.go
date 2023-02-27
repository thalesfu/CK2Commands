package yolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃德瓦EdvaBarony struct {
	feud.BaseBarony
}

var BEdva埃德瓦 feud.Barony = &埃德瓦EdvaBarony{}

func init() {
    f := BEdva埃德瓦.(*埃德瓦EdvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "edva",
		TitleName: "埃德瓦",
		TitleCode: "b_edva",
	}
}
