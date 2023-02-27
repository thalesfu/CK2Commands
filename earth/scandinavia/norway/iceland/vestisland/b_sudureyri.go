package vestisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叙聚雷里SudureyriBarony struct {
	feud.BaseBarony
}

var BSudureyri叙聚雷里 feud.Barony = &叙聚雷里SudureyriBarony{}

func init() {
    f := BSudureyri叙聚雷里.(*叙聚雷里SudureyriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sudureyri",
		TitleName: "叙聚雷里",
		TitleCode: "b_sudureyri",
	}
}
