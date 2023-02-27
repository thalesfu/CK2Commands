package orsha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉钦TalachynBarony struct {
	feud.BaseBarony
}

var BTalachyn塔拉钦 feud.Barony = &塔拉钦TalachynBarony{}

func init() {
    f := BTalachyn塔拉钦.(*塔拉钦TalachynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talachyn",
		TitleName: "塔拉钦",
		TitleCode: "b_talachyn",
	}
}
