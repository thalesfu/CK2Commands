package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔法赫WolfachBarony struct {
	feud.BaseBarony
}

var BWolfach沃尔法赫 feud.Barony = &沃尔法赫WolfachBarony{}

func init() {
	f := BWolfach沃尔法赫.(*沃尔法赫WolfachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wolfach",
		TitleName: "沃尔法赫",
		TitleCode: "b_wolfach",
	}
}
