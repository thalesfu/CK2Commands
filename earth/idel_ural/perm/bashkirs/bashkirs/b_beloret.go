package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别洛列特BeloretBarony struct {
	feud.BaseBarony
}

var BBeloret别洛列特 feud.Barony = &别洛列特BeloretBarony{}

func init() {
    f := BBeloret别洛列特.(*别洛列特BeloretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beloret",
		TitleName: "别洛列特",
		TitleCode: "b_beloret",
	}
}
