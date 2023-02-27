package katehar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内多尔NehtaurBarony struct {
	feud.BaseBarony
}

var BNehtaur内多尔 feud.Barony = &内多尔NehtaurBarony{}

func init() {
    f := BNehtaur内多尔.(*内多尔NehtaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nehtaur",
		TitleName: "内多尔",
		TitleCode: "b_nehtaur",
	}
}
