package lutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列索普尼察亚PeresopnytsiaBarony struct {
	feud.BaseBarony
}

var BPeresopnytsia佩列索普尼察亚 feud.Barony = &佩列索普尼察亚PeresopnytsiaBarony{}

func init() {
    f := BPeresopnytsia佩列索普尼察亚.(*佩列索普尼察亚PeresopnytsiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peresopnytsia",
		TitleName: "佩列索普尼察亚",
		TitleCode: "b_peresopnytsia",
	}
}
