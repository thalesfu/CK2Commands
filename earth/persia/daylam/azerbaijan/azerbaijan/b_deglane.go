package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德格拉尼DeglaneBarony struct {
	feud.BaseBarony
}

var BDeglane德格拉尼 feud.Barony = &德格拉尼DeglaneBarony{}

func init() {
    f := BDeglane德格拉尼.(*德格拉尼DeglaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deglane",
		TitleName: "德格拉尼",
		TitleCode: "b_deglane",
	}
}
