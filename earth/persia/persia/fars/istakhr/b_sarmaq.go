package istakhr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔马格SarmaqBarony struct {
	feud.BaseBarony
}

var BSarmaq萨尔马格 feud.Barony = &萨尔马格SarmaqBarony{}

func init() {
	f := BSarmaq萨尔马格.(*萨尔马格SarmaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarmaq",
		TitleName: "萨尔马格",
		TitleCode: "b_sarmaq",
	}
}
