package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴塔BataBarony struct {
	feud.BaseBarony
}

var BBata巴塔 feud.Barony = &巴塔BataBarony{}

func init() {
    f := BBata巴塔.(*巴塔BataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bata",
		TitleName: "巴塔",
		TitleCode: "b_bata",
	}
}
