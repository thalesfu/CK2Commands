package zemigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶尔加瓦JelgawaBarony struct {
	feud.BaseBarony
}

var BJelgawa叶尔加瓦 feud.Barony = &叶尔加瓦JelgawaBarony{}

func init() {
    f := BJelgawa叶尔加瓦.(*叶尔加瓦JelgawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jelgawa",
		TitleName: "叶尔加瓦",
		TitleCode: "b_jelgawa",
	}
}
