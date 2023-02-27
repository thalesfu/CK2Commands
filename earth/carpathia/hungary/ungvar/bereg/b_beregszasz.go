package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜赖格萨斯BeregszaszBarony struct {
	feud.BaseBarony
}

var BBeregszasz拜赖格萨斯 feud.Barony = &拜赖格萨斯BeregszaszBarony{}

func init() {
    f := BBeregszasz拜赖格萨斯.(*拜赖格萨斯BeregszaszBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beregszasz",
		TitleName: "拜赖格萨斯",
		TitleCode: "b_beregszasz",
	}
}
