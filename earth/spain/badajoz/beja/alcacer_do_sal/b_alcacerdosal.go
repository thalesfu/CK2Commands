package alcacer_do_sal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔堡AlcacerdosalBarony struct {
	feud.BaseBarony
}

var BAlcacerdosal萨尔堡 feud.Barony = &萨尔堡AlcacerdosalBarony{}

func init() {
    f := BAlcacerdosal萨尔堡.(*萨尔堡AlcacerdosalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcacerdosal",
		TitleName: "萨尔堡",
		TitleCode: "b_alcacerdosal",
	}
}
