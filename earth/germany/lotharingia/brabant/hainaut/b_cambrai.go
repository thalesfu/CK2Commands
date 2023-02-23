package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康布雷CambraiBarony struct {
	feud.BaseBarony
}

var BCambrai康布雷 feud.Barony = &康布雷CambraiBarony{}

func init() {
	f := BCambrai康布雷.(*康布雷CambraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cambrai",
		TitleName: "康布雷",
		TitleCode: "b_cambrai",
	}
}
