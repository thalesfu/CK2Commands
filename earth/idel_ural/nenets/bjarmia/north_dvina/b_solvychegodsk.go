package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索里维切戈茨克SolvychegodskBarony struct {
	feud.BaseBarony
}

var BSolvychegodsk索里维切戈茨克 feud.Barony = &索里维切戈茨克SolvychegodskBarony{}

func init() {
    f := BSolvychegodsk索里维切戈茨克.(*索里维切戈茨克SolvychegodskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solvychegodsk",
		TitleName: "索里维切戈茨克",
		TitleCode: "b_solvychegodsk",
	}
}
