package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈特杰达冷HattfjelldalenBarony struct {
	feud.BaseBarony
}

var BHattfjelldalen哈特杰达冷 feud.Barony = &哈特杰达冷HattfjelldalenBarony{}

func init() {
	f := BHattfjelldalen哈特杰达冷.(*哈特杰达冷HattfjelldalenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hattfjelldalen",
		TitleName: "哈特杰达冷",
		TitleCode: "b_hattfjelldalen",
	}
}
