package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼济扎马拉ManziljamalBarony struct {
	feud.BaseBarony
}

var BManziljamal曼济扎马拉 feud.Barony = &曼济扎马拉ManziljamalBarony{}

func init() {
    f := BManziljamal曼济扎马拉.(*曼济扎马拉ManziljamalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manziljamal",
		TitleName: "曼济扎马拉",
		TitleCode: "b_manziljamal",
	}
}
