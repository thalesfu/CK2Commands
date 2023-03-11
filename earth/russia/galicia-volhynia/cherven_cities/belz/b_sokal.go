package belz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索卡尔SokalBarony struct {
	feud.BaseBarony
}

var BSokal索卡尔 feud.Barony = &索卡尔SokalBarony{}

func init() {
    f := BSokal索卡尔.(*索卡尔SokalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sokal",
		TitleName: "索卡尔",
		TitleCode: "b_sokal",
	}
}
