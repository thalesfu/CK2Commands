package skara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶拉克维斯特GalakvistBarony struct {
	feud.BaseBarony
}

var BGalakvist耶拉克维斯特 feud.Barony = &耶拉克维斯特GalakvistBarony{}

func init() {
    f := BGalakvist耶拉克维斯特.(*耶拉克维斯特GalakvistBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galakvist",
		TitleName: "耶拉克维斯特",
		TitleCode: "b_galakvist",
	}
}
