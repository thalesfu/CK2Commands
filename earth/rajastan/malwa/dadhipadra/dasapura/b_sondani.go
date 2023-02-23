package dasapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孙陀尼SondaniBarony struct {
	feud.BaseBarony
}

var BSondani孙陀尼 feud.Barony = &孙陀尼SondaniBarony{}

func init() {
	f := BSondani孙陀尼.(*孙陀尼SondaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sondani",
		TitleName: "孙陀尼",
		TitleCode: "b_sondani",
	}
}
