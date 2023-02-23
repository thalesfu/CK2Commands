package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪奥达摩DieudamourBarony struct {
	feud.BaseBarony
}

var BDieudamour迪奥达摩 feud.Barony = &迪奥达摩DieudamourBarony{}

func init() {
	f := BDieudamour迪奥达摩.(*迪奥达摩DieudamourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dieudamour",
		TitleName: "迪奥达摩",
		TitleCode: "b_dieudamour",
	}
}
