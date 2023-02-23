package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩雷西PerrecyBarony struct {
	feud.BaseBarony
}

var BPerrecy佩雷西 feud.Barony = &佩雷西PerrecyBarony{}

func init() {
	f := BPerrecy佩雷西.(*佩雷西PerrecyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perrecy",
		TitleName: "佩雷西",
		TitleCode: "b_perrecy",
	}
}
