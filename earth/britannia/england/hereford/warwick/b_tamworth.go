package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔姆沃思TamworthBarony struct {
	feud.BaseBarony
}

var BTamworth塔姆沃思 feud.Barony = &塔姆沃思TamworthBarony{}

func init() {
    f := BTamworth塔姆沃思.(*塔姆沃思TamworthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamworth",
		TitleName: "塔姆沃思",
		TitleCode: "b_tamworth",
	}
}
