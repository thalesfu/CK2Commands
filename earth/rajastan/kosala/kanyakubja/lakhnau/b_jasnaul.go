package lakhnau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶娑那乌罗JasnaulBarony struct {
	feud.BaseBarony
}

var BJasnaul耶娑那乌罗 feud.Barony = &耶娑那乌罗JasnaulBarony{}

func init() {
	f := BJasnaul耶娑那乌罗.(*耶娑那乌罗JasnaulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jasnaul",
		TitleName: "耶娑那乌罗",
		TitleCode: "b_jasnaul",
	}
}
