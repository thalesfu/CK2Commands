package smolensk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格涅兹多沃GnyozdovoBarony struct {
	feud.BaseBarony
}

var BGnyozdovo格涅兹多沃 feud.Barony = &格涅兹多沃GnyozdovoBarony{}

func init() {
	f := BGnyozdovo格涅兹多沃.(*格涅兹多沃GnyozdovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gnyozdovo",
		TitleName: "格涅兹多沃",
		TitleCode: "b_gnyozdovo",
	}
}
