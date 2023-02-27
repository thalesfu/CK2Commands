package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔曼TamanBarony struct {
	feud.BaseBarony
}

var BTaman塔曼 feud.Barony = &塔曼TamanBarony{}

func init() {
    f := BTaman塔曼.(*塔曼TamanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taman",
		TitleName: "塔曼",
		TitleCode: "b_taman",
	}
}
