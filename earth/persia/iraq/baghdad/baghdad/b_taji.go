package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔吉TajiBarony struct {
	feud.BaseBarony
}

var BTaji塔吉 feud.Barony = &塔吉TajiBarony{}

func init() {
    f := BTaji塔吉.(*塔吉TajiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taji",
		TitleName: "塔吉",
		TitleCode: "b_taji",
	}
}
