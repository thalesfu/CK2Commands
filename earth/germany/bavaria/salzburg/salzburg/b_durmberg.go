package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪姆贝格DurmbergBarony struct {
	feud.BaseBarony
}

var BDurmberg迪姆贝格 feud.Barony = &迪姆贝格DurmbergBarony{}

func init() {
    f := BDurmberg迪姆贝格.(*迪姆贝格DurmbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durmberg",
		TitleName: "迪姆贝格",
		TitleCode: "b_durmberg",
	}
}
