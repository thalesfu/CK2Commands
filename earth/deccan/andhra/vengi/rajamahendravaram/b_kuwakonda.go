package rajamahendravaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠婆军荼KuwakondaBarony struct {
	feud.BaseBarony
}

var BKuwakonda鸠婆军荼 feud.Barony = &鸠婆军荼KuwakondaBarony{}

func init() {
    f := BKuwakonda鸠婆军荼.(*鸠婆军荼KuwakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuwakonda",
		TitleName: "鸠婆军荼",
		TitleCode: "b_kuwakonda",
	}
}
