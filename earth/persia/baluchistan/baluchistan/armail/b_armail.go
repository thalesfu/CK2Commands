package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚梅尔ArmailBarony struct {
	feud.BaseBarony
}

var BArmail亚梅尔 feud.Barony = &亚梅尔ArmailBarony{}

func init() {
    f := BArmail亚梅尔.(*亚梅尔ArmailBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "armail",
		TitleName: "亚梅尔",
		TitleCode: "b_armail",
	}
}
