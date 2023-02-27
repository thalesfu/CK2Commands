package amorion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲洛梅利翁PhilomelionBarony struct {
	feud.BaseBarony
}

var BPhilomelion菲洛梅利翁 feud.Barony = &菲洛梅利翁PhilomelionBarony{}

func init() {
    f := BPhilomelion菲洛梅利翁.(*菲洛梅利翁PhilomelionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "philomelion",
		TitleName: "菲洛梅利翁",
		TitleCode: "b_philomelion",
	}
}
