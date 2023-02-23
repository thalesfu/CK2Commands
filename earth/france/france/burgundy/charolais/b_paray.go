package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕雷ParayBarony struct {
	feud.BaseBarony
}

var BParay帕雷 feud.Barony = &帕雷ParayBarony{}

func init() {
	f := BParay帕雷.(*帕雷ParayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paray",
		TitleName: "帕雷",
		TitleCode: "b_paray",
	}
}
