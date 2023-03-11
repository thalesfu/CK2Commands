package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切斯特勒街Chester_le_streetBarony struct {
	feud.BaseBarony
}

var BChester_le_street切斯特勒街 feud.Barony = &切斯特勒街Chester_le_streetBarony{}

func init() {
    f := BChester_le_street切斯特勒街.(*切斯特勒街Chester_le_streetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chester_le_street",
		TitleName: "切斯特勒街",
		TitleCode: "b_chester-le-street",
	}
}
