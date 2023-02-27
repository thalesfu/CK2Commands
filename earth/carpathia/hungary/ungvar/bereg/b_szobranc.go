package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索布兰茨SzobrancBarony struct {
	feud.BaseBarony
}

var BSzobranc索布兰茨 feud.Barony = &索布兰茨SzobrancBarony{}

func init() {
    f := BSzobranc索布兰茨.(*索布兰茨SzobrancBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szobranc",
		TitleName: "索布兰茨",
		TitleCode: "b_szobranc",
	}
}
