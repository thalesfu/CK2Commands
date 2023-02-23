package skara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福什霍尔姆ForsholmBarony struct {
	feud.BaseBarony
}

var BForsholm福什霍尔姆 feud.Barony = &福什霍尔姆ForsholmBarony{}

func init() {
	f := BForsholm福什霍尔姆.(*福什霍尔姆ForsholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forsholm",
		TitleName: "福什霍尔姆",
		TitleCode: "b_forsholm",
	}
}
