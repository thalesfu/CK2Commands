package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米德尔堡MiddelburgBarony struct {
	feud.BaseBarony
}

var BMiddelburg米德尔堡 feud.Barony = &米德尔堡MiddelburgBarony{}

func init() {
    f := BMiddelburg米德尔堡.(*米德尔堡MiddelburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "middelburg",
		TitleName: "米德尔堡",
		TitleCode: "b_middelburg",
	}
}
