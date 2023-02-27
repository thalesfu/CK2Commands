package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毾㲪山KuhtaftanBarony struct {
	feud.BaseBarony
}

var BKuhtaftan毾㲪山 feud.Barony = &毾㲪山KuhtaftanBarony{}

func init() {
    f := BKuhtaftan毾㲪山.(*毾㲪山KuhtaftanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuhtaftan",
		TitleName: "毾㲪山",
		TitleCode: "b_kuhtaftan",
	}
}
