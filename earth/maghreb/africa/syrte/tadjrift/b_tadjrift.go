package tadjrift

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔杰里夫特TadjriftBarony struct {
	feud.BaseBarony
}

var BTadjrift塔杰里夫特 feud.Barony = &塔杰里夫特TadjriftBarony{}

func init() {
    f := BTadjrift塔杰里夫特.(*塔杰里夫特TadjriftBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tadjrift",
		TitleName: "塔杰里夫特",
		TitleCode: "b_tadjrift",
	}
}
