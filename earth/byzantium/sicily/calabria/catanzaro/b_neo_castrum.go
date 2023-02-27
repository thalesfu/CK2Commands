package catanzaro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新堡Neo_castrumBarony struct {
	feud.BaseBarony
}

var BNeo_castrum新堡 feud.Barony = &新堡Neo_castrumBarony{}

func init() {
    f := BNeo_castrum新堡.(*新堡Neo_castrumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neo_castrum",
		TitleName: "新堡",
		TitleCode: "b_neo_castrum",
	}
}
