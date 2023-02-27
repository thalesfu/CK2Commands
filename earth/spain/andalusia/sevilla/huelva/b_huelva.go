package huelva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔瓦HuelvaBarony struct {
	feud.BaseBarony
}

var BHuelva韦尔瓦 feud.Barony = &韦尔瓦HuelvaBarony{}

func init() {
    f := BHuelva韦尔瓦.(*韦尔瓦HuelvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huelva",
		TitleName: "韦尔瓦",
		TitleCode: "b_huelva",
	}
}
