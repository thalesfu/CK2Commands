package alcacer_do_sal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦迪堡CastroverdeBarony struct {
	feud.BaseBarony
}

var BCastroverde韦迪堡 feud.Barony = &韦迪堡CastroverdeBarony{}

func init() {
    f := BCastroverde韦迪堡.(*韦迪堡CastroverdeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castroverde",
		TitleName: "韦迪堡",
		TitleCode: "b_castroverde",
	}
}
