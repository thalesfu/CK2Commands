package kartaly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉普克尔TalapkerBarony struct {
	feud.BaseBarony
}

var BTalapker塔拉普克尔 feud.Barony = &塔拉普克尔TalapkerBarony{}

func init() {
	f := BTalapker塔拉普克尔.(*塔拉普克尔TalapkerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talapker",
		TitleName: "塔拉普克尔",
		TitleCode: "b_talapker",
	}
}
