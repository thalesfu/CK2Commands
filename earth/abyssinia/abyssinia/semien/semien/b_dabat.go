package semien

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达巴特DabatBarony struct {
	feud.BaseBarony
}

var BDabat达巴特 feud.Barony = &达巴特DabatBarony{}

func init() {
	f := BDabat达巴特.(*达巴特DabatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dabat",
		TitleName: "达巴特",
		TitleCode: "b_dabat",
	}
}
