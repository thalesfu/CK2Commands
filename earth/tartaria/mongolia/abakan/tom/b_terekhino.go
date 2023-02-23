package tom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷列希诺TerekhinoBarony struct {
	feud.BaseBarony
}

var BTerekhino捷列希诺 feud.Barony = &捷列希诺TerekhinoBarony{}

func init() {
	f := BTerekhino捷列希诺.(*捷列希诺TerekhinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terekhino",
		TitleName: "捷列希诺",
		TitleCode: "b_terekhino",
	}
}
