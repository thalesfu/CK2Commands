package tadmekka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔得迈卡TadmekkaBarony struct {
	feud.BaseBarony
}

var BTadmekka塔得迈卡 feud.Barony = &塔得迈卡TadmekkaBarony{}

func init() {
    f := BTadmekka塔得迈卡.(*塔得迈卡TadmekkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tadmekka",
		TitleName: "塔得迈卡",
		TitleCode: "b_tadmekka",
	}
}
