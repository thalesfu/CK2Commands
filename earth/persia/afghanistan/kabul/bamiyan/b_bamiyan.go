package bamiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梵衍那BamiyanBarony struct {
	feud.BaseBarony
}

var BBamiyan梵衍那 feud.Barony = &梵衍那BamiyanBarony{}

func init() {
    f := BBamiyan梵衍那.(*梵衍那BamiyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bamiyan",
		TitleName: "梵衍那",
		TitleCode: "b_bamiyan",
	}
}
