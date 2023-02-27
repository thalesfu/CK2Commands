package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏格阿赫SuqashBarony struct {
	feud.BaseBarony
}

var BSuqash苏格阿赫 feud.Barony = &苏格阿赫SuqashBarony{}

func init() {
    f := BSuqash苏格阿赫.(*苏格阿赫SuqashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suqash",
		TitleName: "苏格阿赫",
		TitleCode: "b_suqash",
	}
}
