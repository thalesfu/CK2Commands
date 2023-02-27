package kassala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔塞特GhersetBarony struct {
	feud.BaseBarony
}

var BGherset戈尔塞特 feud.Barony = &戈尔塞特GhersetBarony{}

func init() {
    f := BGherset戈尔塞特.(*戈尔塞特GhersetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gherset",
		TitleName: "戈尔塞特",
		TitleCode: "b_gherset",
	}
}
