package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴士拉BasraBarony struct {
	feud.BaseBarony
}

var BBasra巴士拉 feud.Barony = &巴士拉BasraBarony{}

func init() {
    f := BBasra巴士拉.(*巴士拉BasraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basra",
		TitleName: "巴士拉",
		TitleCode: "b_basra",
	}
}
