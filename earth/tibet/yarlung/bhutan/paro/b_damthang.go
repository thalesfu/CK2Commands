package paro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 当塘DamthangBarony struct {
	feud.BaseBarony
}

var BDamthang当塘 feud.Barony = &当塘DamthangBarony{}

func init() {
    f := BDamthang当塘.(*当塘DamthangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damthang",
		TitleName: "当塘",
		TitleCode: "b_damthang",
	}
}
