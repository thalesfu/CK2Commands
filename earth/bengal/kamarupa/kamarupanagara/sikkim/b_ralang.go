package sikkim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉浪RalangBarony struct {
	feud.BaseBarony
}

var BRalang拉浪 feud.Barony = &拉浪RalangBarony{}

func init() {
    f := BRalang拉浪.(*拉浪RalangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ralang",
		TitleName: "拉浪",
		TitleCode: "b_ralang",
	}
}
