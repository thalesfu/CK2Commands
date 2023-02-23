package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛巴利亚CimballaBarony struct {
	feud.BaseBarony
}

var BCimballa辛巴利亚 feud.Barony = &辛巴利亚CimballaBarony{}

func init() {
	f := BCimballa辛巴利亚.(*辛巴利亚CimballaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cimballa",
		TitleName: "辛巴利亚",
		TitleCode: "b_cimballa",
	}
}
