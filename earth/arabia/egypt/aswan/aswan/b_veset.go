package aswan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维瑟特VesetBarony struct {
	feud.BaseBarony
}

var BVeset维瑟特 feud.Barony = &维瑟特VesetBarony{}

func init() {
    f := BVeset维瑟特.(*维瑟特VesetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veset",
		TitleName: "维瑟特",
		TitleCode: "b_veset",
	}
}
