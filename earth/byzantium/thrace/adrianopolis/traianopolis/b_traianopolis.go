package traianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图拉真波利斯TraianopolisBarony struct {
	feud.BaseBarony
}

var BTraianopolis图拉真波利斯 feud.Barony = &图拉真波利斯TraianopolisBarony{}

func init() {
    f := BTraianopolis图拉真波利斯.(*图拉真波利斯TraianopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "traianopolis",
		TitleName: "图拉真波利斯",
		TitleCode: "b_traianopolis",
	}
}
