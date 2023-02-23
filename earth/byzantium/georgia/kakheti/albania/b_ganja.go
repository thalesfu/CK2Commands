package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 占贾GanjaBarony struct {
	feud.BaseBarony
}

var BGanja占贾 feud.Barony = &占贾GanjaBarony{}

func init() {
	f := BGanja占贾.(*占贾GanjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ganja",
		TitleName: "占贾",
		TitleCode: "b_ganja",
	}
}
