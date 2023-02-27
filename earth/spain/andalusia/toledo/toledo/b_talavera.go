package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉韦拉TalaveraBarony struct {
	feud.BaseBarony
}

var BTalavera塔拉韦拉 feud.Barony = &塔拉韦拉TalaveraBarony{}

func init() {
    f := BTalavera塔拉韦拉.(*塔拉韦拉TalaveraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talavera",
		TitleName: "塔拉韦拉",
		TitleCode: "b_talavera",
	}
}
