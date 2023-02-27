package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔达ThattaBarony struct {
	feud.BaseBarony
}

var BThatta塔达 feud.Barony = &塔达ThattaBarony{}

func init() {
    f := BThatta塔达.(*塔达ThattaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thatta",
		TitleName: "塔达",
		TitleCode: "b_thatta",
	}
}
