package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡萨雷斯CasaresBarony struct {
	feud.BaseBarony
}

var BCasares卡萨雷斯 feud.Barony = &卡萨雷斯CasaresBarony{}

func init() {
    f := BCasares卡萨雷斯.(*卡萨雷斯CasaresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "casares",
		TitleName: "卡萨雷斯",
		TitleCode: "b_casares",
	}
}
