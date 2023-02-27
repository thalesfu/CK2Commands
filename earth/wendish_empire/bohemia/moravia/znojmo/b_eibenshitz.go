package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾本希茨EibenshitzBarony struct {
	feud.BaseBarony
}

var BEibenshitz艾本希茨 feud.Barony = &艾本希茨EibenshitzBarony{}

func init() {
    f := BEibenshitz艾本希茨.(*艾本希茨EibenshitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eibenshitz",
		TitleName: "艾本希茨",
		TitleCode: "b_eibenshitz",
	}
}
