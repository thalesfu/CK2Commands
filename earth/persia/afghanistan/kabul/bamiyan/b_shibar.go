package bamiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希巴尔ShibarBarony struct {
	feud.BaseBarony
}

var BShibar希巴尔 feud.Barony = &希巴尔ShibarBarony{}

func init() {
	f := BShibar希巴尔.(*希巴尔ShibarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shibar",
		TitleName: "希巴尔",
		TitleCode: "b_shibar",
	}
}
