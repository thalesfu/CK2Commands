package kashgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埠口BukouBarony struct {
	feud.BaseBarony
}

var BBukou埠口 feud.Barony = &埠口BukouBarony{}

func init() {
    f := BBukou埠口.(*埠口BukouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bukou",
		TitleName: "埠口",
		TitleCode: "b_bukou",
	}
}
