package syrt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨马拉SamaraBarony struct {
	feud.BaseBarony
}

var BSamara萨马拉 feud.Barony = &萨马拉SamaraBarony{}

func init() {
    f := BSamara萨马拉.(*萨马拉SamaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samara",
		TitleName: "萨马拉",
		TitleCode: "b_samara",
	}
}
