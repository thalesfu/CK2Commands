package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔里法TarifaBarony struct {
	feud.BaseBarony
}

var BTarifa塔里法 feud.Barony = &塔里法TarifaBarony{}

func init() {
	f := BTarifa塔里法.(*塔里法TarifaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarifa",
		TitleName: "塔里法",
		TitleCode: "b_tarifa",
	}
}
