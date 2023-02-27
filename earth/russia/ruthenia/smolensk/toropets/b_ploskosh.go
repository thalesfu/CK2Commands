package toropets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普洛斯科希PloskoshBarony struct {
	feud.BaseBarony
}

var BPloskosh普洛斯科希 feud.Barony = &普洛斯科希PloskoshBarony{}

func init() {
    f := BPloskosh普洛斯科希.(*普洛斯科希PloskoshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ploskosh",
		TitleName: "普洛斯科希",
		TitleCode: "b_ploskosh",
	}
}
