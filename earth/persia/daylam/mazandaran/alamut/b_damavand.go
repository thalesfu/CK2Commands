package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达马万德DamavandBarony struct {
	feud.BaseBarony
}

var BDamavand达马万德 feud.Barony = &达马万德DamavandBarony{}

func init() {
	f := BDamavand达马万德.(*达马万德DamavandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damavand",
		TitleName: "达马万德",
		TitleCode: "b_damavand",
	}
}
