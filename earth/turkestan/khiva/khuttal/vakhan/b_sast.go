package vakhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛斯特SastBarony struct {
	feud.BaseBarony
}

var BSast赛斯特 feud.Barony = &赛斯特SastBarony{}

func init() {
	f := BSast赛斯特.(*赛斯特SastBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sast",
		TitleName: "赛斯特",
		TitleCode: "b_sast",
	}
}
