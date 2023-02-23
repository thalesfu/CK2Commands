package miass

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马雷希MalyshiBarony struct {
	feud.BaseBarony
}

var BMalyshi马雷希 feud.Barony = &马雷希MalyshiBarony{}

func init() {
	f := BMalyshi马雷希.(*马雷希MalyshiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malyshi",
		TitleName: "马雷希",
		TitleCode: "b_malyshi",
	}
}
