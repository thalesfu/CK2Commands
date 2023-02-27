package sozopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加多CadiBarony struct {
	feud.BaseBarony
}

var BCadi加多 feud.Barony = &加多CadiBarony{}

func init() {
    f := BCadi加多.(*加多CadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cadi",
		TitleName: "加多",
		TitleCode: "b_cadi",
	}
}
