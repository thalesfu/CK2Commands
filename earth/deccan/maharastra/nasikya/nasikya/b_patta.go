package nasikya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕塔PattaBarony struct {
	feud.BaseBarony
}

var BPatta帕塔 feud.Barony = &帕塔PattaBarony{}

func init() {
    f := BPatta帕塔.(*帕塔PattaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patta",
		TitleName: "帕塔",
		TitleCode: "b_patta",
	}
}
