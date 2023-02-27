package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达玛瓦特DamqwatBarony struct {
	feud.BaseBarony
}

var BDamqwat达玛瓦特 feud.Barony = &达玛瓦特DamqwatBarony{}

func init() {
    f := BDamqwat达玛瓦特.(*达玛瓦特DamqwatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damqwat",
		TitleName: "达玛瓦特",
		TitleCode: "b_damqwat",
	}
}
