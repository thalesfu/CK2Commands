package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯拜HesbaieBarony struct {
	feud.BaseBarony
}

var BHesbaie埃斯拜 feud.Barony = &埃斯拜HesbaieBarony{}

func init() {
    f := BHesbaie埃斯拜.(*埃斯拜HesbaieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hesbaie",
		TitleName: "埃斯拜",
		TitleCode: "b_hesbaie",
	}
}
