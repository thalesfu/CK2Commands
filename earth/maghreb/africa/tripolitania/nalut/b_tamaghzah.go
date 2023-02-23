package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔马格扎TamaghzahBarony struct {
	feud.BaseBarony
}

var BTamaghzah塔马格扎 feud.Barony = &塔马格扎TamaghzahBarony{}

func init() {
	f := BTamaghzah塔马格扎.(*塔马格扎TamaghzahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamaghzah",
		TitleName: "塔马格扎",
		TitleCode: "b_tamaghzah",
	}
}
