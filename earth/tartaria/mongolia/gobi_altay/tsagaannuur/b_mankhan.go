package tsagaannuur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼汗MankhanBarony struct {
	feud.BaseBarony
}

var BMankhan曼汗 feud.Barony = &曼汗MankhanBarony{}

func init() {
    f := BMankhan曼汗.(*曼汗MankhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mankhan",
		TitleName: "曼汗",
		TitleCode: "b_mankhan",
	}
}
