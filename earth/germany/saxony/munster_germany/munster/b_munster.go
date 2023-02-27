package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明斯特MunsterBarony struct {
	feud.BaseBarony
}

var BMunster明斯特 feud.Barony = &明斯特MunsterBarony{}

func init() {
    f := BMunster明斯特.(*明斯特MunsterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "munster",
		TitleName: "明斯特",
		TitleCode: "b_munster",
	}
}
