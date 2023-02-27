package mountain_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维博尔纳亚VybornayaBarony struct {
	feud.BaseBarony
}

var BVybornaya维博尔纳亚 feud.Barony = &维博尔纳亚VybornayaBarony{}

func init() {
    f := BVybornaya维博尔纳亚.(*维博尔纳亚VybornayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vybornaya",
		TitleName: "维博尔纳亚",
		TitleCode: "b_vybornaya",
	}
}
