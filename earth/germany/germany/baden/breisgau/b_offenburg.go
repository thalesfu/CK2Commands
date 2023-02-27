package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥芬堡OffenburgBarony struct {
	feud.BaseBarony
}

var BOffenburg奥芬堡 feud.Barony = &奥芬堡OffenburgBarony{}

func init() {
    f := BOffenburg奥芬堡.(*奥芬堡OffenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "offenburg",
		TitleName: "奥芬堡",
		TitleCode: "b_offenburg",
	}
}
