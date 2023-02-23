package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔雷加TarregaBarony struct {
	feud.BaseBarony
}

var BTarrega塔雷加 feud.Barony = &塔雷加TarregaBarony{}

func init() {
	f := BTarrega塔雷加.(*塔雷加TarregaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarrega",
		TitleName: "塔雷加",
		TitleCode: "b_tarrega",
	}
}
