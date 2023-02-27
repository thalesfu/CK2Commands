package nizhny_novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴甫洛沃PavlovoBarony struct {
	feud.BaseBarony
}

var BPavlovo巴甫洛沃 feud.Barony = &巴甫洛沃PavlovoBarony{}

func init() {
    f := BPavlovo巴甫洛沃.(*巴甫洛沃PavlovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pavlovo",
		TitleName: "巴甫洛沃",
		TitleCode: "b_pavlovo",
	}
}
