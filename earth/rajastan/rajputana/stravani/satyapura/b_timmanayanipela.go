package satyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顶摩尼那吠罗TimmanayanipelaBarony struct {
	feud.BaseBarony
}

var BTimmanayanipela顶摩尼那吠罗 feud.Barony = &顶摩尼那吠罗TimmanayanipelaBarony{}

func init() {
    f := BTimmanayanipela顶摩尼那吠罗.(*顶摩尼那吠罗TimmanayanipelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timmanayanipela",
		TitleName: "顶摩尼那吠罗",
		TitleCode: "b_timmanayanipela",
	}
}
