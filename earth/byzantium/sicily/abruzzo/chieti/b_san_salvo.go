package chieti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣萨尔沃San_salvoBarony struct {
	feud.BaseBarony
}

var BSan_salvo圣萨尔沃 feud.Barony = &圣萨尔沃San_salvoBarony{}

func init() {
    f := BSan_salvo圣萨尔沃.(*圣萨尔沃San_salvoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "san_salvo",
		TitleName: "圣萨尔沃",
		TitleCode: "b_san_salvo",
	}
}
