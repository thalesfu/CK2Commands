package kalaus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴佐维BazovyyBarony struct {
	feud.BaseBarony
}

var BBazovyy巴佐维 feud.Barony = &巴佐维BazovyyBarony{}

func init() {
    f := BBazovyy巴佐维.(*巴佐维BazovyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bazovyy",
		TitleName: "巴佐维",
		TitleCode: "b_bazovyy",
	}
}
