package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃瓦尔EibarBarony struct {
	feud.BaseBarony
}

var BEibar埃瓦尔 feud.Barony = &埃瓦尔EibarBarony{}

func init() {
    f := BEibar埃瓦尔.(*埃瓦尔EibarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eibar",
		TitleName: "埃瓦尔",
		TitleCode: "b_eibar",
	}
}
