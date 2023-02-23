package aksum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿迪格拉特AdigratBarony struct {
	feud.BaseBarony
}

var BAdigrat阿迪格拉特 feud.Barony = &阿迪格拉特AdigratBarony{}

func init() {
	f := BAdigrat阿迪格拉特.(*阿迪格拉特AdigratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adigrat",
		TitleName: "阿迪格拉特",
		TitleCode: "b_adigrat",
	}
}
