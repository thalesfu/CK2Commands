package rama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯雷布雷尼察SrebrenicaBarony struct {
	feud.BaseBarony
}

var BSrebrenica斯雷布雷尼察 feud.Barony = &斯雷布雷尼察SrebrenicaBarony{}

func init() {
    f := BSrebrenica斯雷布雷尼察.(*斯雷布雷尼察SrebrenicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srebrenica",
		TitleName: "斯雷布雷尼察",
		TitleCode: "b_srebrenica",
	}
}
