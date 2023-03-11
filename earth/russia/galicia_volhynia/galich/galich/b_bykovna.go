package galich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布科维纳BykovnaBarony struct {
	feud.BaseBarony
}

var BBykovna布科维纳 feud.Barony = &布科维纳BykovnaBarony{}

func init() {
    f := BBykovna布科维纳.(*布科维纳BykovnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bykovna",
		TitleName: "布科维纳",
		TitleCode: "b_bykovna",
	}
}
