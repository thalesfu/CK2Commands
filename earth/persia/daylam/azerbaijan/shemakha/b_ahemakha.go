package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍马哈AhemakhaBarony struct {
	feud.BaseBarony
}

var BAhemakha舍马哈 feud.Barony = &舍马哈AhemakhaBarony{}

func init() {
    f := BAhemakha舍马哈.(*舍马哈AhemakhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahemakha",
		TitleName: "舍马哈",
		TitleCode: "b_ahemakha",
	}
}
