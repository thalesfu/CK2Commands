package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特格堡StegeborgBarony struct {
	feud.BaseBarony
}

var BStegeborg斯特格堡 feud.Barony = &斯特格堡StegeborgBarony{}

func init() {
    f := BStegeborg斯特格堡.(*斯特格堡StegeborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stegeborg",
		TitleName: "斯特格堡",
		TitleCode: "b_stegeborg",
	}
}
