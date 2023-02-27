package sakmara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥伦堡OrenburgBarony struct {
	feud.BaseBarony
}

var BOrenburg奥伦堡 feud.Barony = &奥伦堡OrenburgBarony{}

func init() {
    f := BOrenburg奥伦堡.(*奥伦堡OrenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orenburg",
		TitleName: "奥伦堡",
		TitleCode: "b_orenburg",
	}
}
