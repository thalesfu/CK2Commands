package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉巴TarabaBarony struct {
	feud.BaseBarony
}

var BTaraba塔拉巴 feud.Barony = &塔拉巴TarabaBarony{}

func init() {
	f := BTaraba塔拉巴.(*塔拉巴TarabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taraba",
		TitleName: "塔拉巴",
		TitleCode: "b_taraba",
	}
}
