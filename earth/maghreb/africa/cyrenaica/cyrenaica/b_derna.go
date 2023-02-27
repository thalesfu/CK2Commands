package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德尔纳DernaBarony struct {
	feud.BaseBarony
}

var BDerna德尔纳 feud.Barony = &德尔纳DernaBarony{}

func init() {
    f := BDerna德尔纳.(*德尔纳DernaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derna",
		TitleName: "德尔纳",
		TitleCode: "b_derna",
	}
}
