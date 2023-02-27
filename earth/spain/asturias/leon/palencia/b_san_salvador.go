package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣萨尔瓦多San_salvadorBarony struct {
	feud.BaseBarony
}

var BSan_salvador圣萨尔瓦多 feud.Barony = &圣萨尔瓦多San_salvadorBarony{}

func init() {
    f := BSan_salvador圣萨尔瓦多.(*圣萨尔瓦多San_salvadorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "san_salvador",
		TitleName: "圣萨尔瓦多",
		TitleCode: "b_san_salvador",
	}
}
