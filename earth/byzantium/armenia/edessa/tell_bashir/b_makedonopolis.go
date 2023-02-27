package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马其顿诺波利斯MakedonopolisBarony struct {
	feud.BaseBarony
}

var BMakedonopolis马其顿诺波利斯 feud.Barony = &马其顿诺波利斯MakedonopolisBarony{}

func init() {
    f := BMakedonopolis马其顿诺波利斯.(*马其顿诺波利斯MakedonopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makedonopolis",
		TitleName: "马其顿诺波利斯",
		TitleCode: "b_makedonopolis",
	}
}
