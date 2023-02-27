package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏曼纳亚SumanayBarony struct {
	feud.BaseBarony
}

var BSumanay苏曼纳亚 feud.Barony = &苏曼纳亚SumanayBarony{}

func init() {
    f := BSumanay苏曼纳亚.(*苏曼纳亚SumanayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumanay",
		TitleName: "苏曼纳亚",
		TitleCode: "b_sumanay",
	}
}
