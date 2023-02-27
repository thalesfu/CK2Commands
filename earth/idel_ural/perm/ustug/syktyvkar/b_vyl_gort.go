package syktyvkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维利戈尔特Vyl_gortBarony struct {
	feud.BaseBarony
}

var BVyl_gort维利戈尔特 feud.Barony = &维利戈尔特Vyl_gortBarony{}

func init() {
    f := BVyl_gort维利戈尔特.(*维利戈尔特Vyl_gortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyl_gort",
		TitleName: "维利戈尔特",
		TitleCode: "b_vyl_gort",
	}
}
