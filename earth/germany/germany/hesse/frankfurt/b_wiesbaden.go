package frankfurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威斯巴登WiesbadenBarony struct {
	feud.BaseBarony
}

var BWiesbaden威斯巴登 feud.Barony = &威斯巴登WiesbadenBarony{}

func init() {
	f := BWiesbaden威斯巴登.(*威斯巴登WiesbadenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wiesbaden",
		TitleName: "威斯巴登",
		TitleCode: "b_wiesbaden",
	}
}
