package neumark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉瓦DrawaBarony struct {
	feud.BaseBarony
}

var BDrawa德拉瓦 feud.Barony = &德拉瓦DrawaBarony{}

func init() {
    f := BDrawa德拉瓦.(*德拉瓦DrawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drawa",
		TitleName: "德拉瓦",
		TitleCode: "b_drawa",
	}
}
