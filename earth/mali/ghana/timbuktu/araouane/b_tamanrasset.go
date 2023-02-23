package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔曼拉塞特TamanrassetBarony struct {
	feud.BaseBarony
}

var BTamanrasset塔曼拉塞特 feud.Barony = &塔曼拉塞特TamanrassetBarony{}

func init() {
	f := BTamanrasset塔曼拉塞特.(*塔曼拉塞特TamanrassetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamanrasset",
		TitleName: "塔曼拉塞特",
		TitleCode: "b_tamanrasset",
	}
}
