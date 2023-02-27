package kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库斯马KusmaBarony struct {
	feud.BaseBarony
}

var BKusma库斯马 feud.Barony = &库斯马KusmaBarony{}

func init() {
    f := BKusma库斯马.(*库斯马KusmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kusma",
		TitleName: "库斯马",
		TitleCode: "b_kusma",
	}
}
