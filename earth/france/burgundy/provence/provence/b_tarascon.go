package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉斯孔TarasconBarony struct {
	feud.BaseBarony
}

var BTarascon塔拉斯孔 feud.Barony = &塔拉斯孔TarasconBarony{}

func init() {
    f := BTarascon塔拉斯孔.(*塔拉斯孔TarasconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarascon",
		TitleName: "塔拉斯孔",
		TitleCode: "b_tarascon",
	}
}
