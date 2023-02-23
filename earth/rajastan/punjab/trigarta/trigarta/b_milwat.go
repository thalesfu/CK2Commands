package trigarta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 密瓦特MilwatBarony struct {
	feud.BaseBarony
}

var BMilwat密瓦特 feud.Barony = &密瓦特MilwatBarony{}

func init() {
	f := BMilwat密瓦特.(*密瓦特MilwatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "milwat",
		TitleName: "密瓦特",
		TitleCode: "b_milwat",
	}
}
