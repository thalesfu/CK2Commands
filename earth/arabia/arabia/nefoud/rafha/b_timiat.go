package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提米特TimiatBarony struct {
	feud.BaseBarony
}

var BTimiat提米特 feud.Barony = &提米特TimiatBarony{}

func init() {
    f := BTimiat提米特.(*提米特TimiatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timiat",
		TitleName: "提米特",
		TitleCode: "b_timiat",
	}
}
