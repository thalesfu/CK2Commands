package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大沙尔蒙Grand_charmontBarony struct {
	feud.BaseBarony
}

var BGrand_charmont大沙尔蒙 feud.Barony = &大沙尔蒙Grand_charmontBarony{}

func init() {
    f := BGrand_charmont大沙尔蒙.(*大沙尔蒙Grand_charmontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grand_charmont",
		TitleName: "大沙尔蒙",
		TitleCode: "b_grand_charmont",
	}
}
