package roslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武吉姆斯克VugimskBarony struct {
	feud.BaseBarony
}

var BVugimsk武吉姆斯克 feud.Barony = &武吉姆斯克VugimskBarony{}

func init() {
    f := BVugimsk武吉姆斯克.(*武吉姆斯克VugimskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vugimsk",
		TitleName: "武吉姆斯克",
		TitleCode: "b_vugimsk",
	}
}
