package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔布TabouBarony struct {
	feud.BaseBarony
}

var BTabou塔布 feud.Barony = &塔布TabouBarony{}

func init() {
	f := BTabou塔布.(*塔布TabouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tabou",
		TitleName: "塔布",
		TitleCode: "b_tabou",
	}
}
