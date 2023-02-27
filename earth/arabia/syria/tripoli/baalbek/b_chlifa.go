package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什里法ChlifaBarony struct {
	feud.BaseBarony
}

var BChlifa什里法 feud.Barony = &什里法ChlifaBarony{}

func init() {
    f := BChlifa什里法.(*什里法ChlifaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chlifa",
		TitleName: "什里法",
		TitleCode: "b_chlifa",
	}
}
