package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔埃巴TaebahBarony struct {
	feud.BaseBarony
}

var BTaebah塔埃巴 feud.Barony = &塔埃巴TaebahBarony{}

func init() {
    f := BTaebah塔埃巴.(*塔埃巴TaebahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taebah",
		TitleName: "塔埃巴",
		TitleCode: "b_taebah",
	}
}
