package freistadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格赖恩GreinBarony struct {
	feud.BaseBarony
}

var BGrein格赖恩 feud.Barony = &格赖恩GreinBarony{}

func init() {
	f := BGrein格赖恩.(*格赖恩GreinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grein",
		TitleName: "格赖恩",
		TitleCode: "b_grein",
	}
}
