package kalyani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨娑罗SarsalaBarony struct {
	feud.BaseBarony
}

var BSarsala萨娑罗 feud.Barony = &萨娑罗SarsalaBarony{}

func init() {
	f := BSarsala萨娑罗.(*萨娑罗SarsalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarsala",
		TitleName: "萨娑罗",
		TitleCode: "b_sarsala",
	}
}
