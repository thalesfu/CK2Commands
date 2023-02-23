package laodikeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛阿斯RhoasBarony struct {
	feud.BaseBarony
}

var BRhoas洛阿斯 feud.Barony = &洛阿斯RhoasBarony{}

func init() {
	f := BRhoas洛阿斯.(*洛阿斯RhoasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rhoas",
		TitleName: "洛阿斯",
		TitleCode: "b_rhoas",
	}
}
