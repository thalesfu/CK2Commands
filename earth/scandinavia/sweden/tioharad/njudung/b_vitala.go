package njudung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维塔拉VitalaBarony struct {
	feud.BaseBarony
}

var BVitala维塔拉 feud.Barony = &维塔拉VitalaBarony{}

func init() {
    f := BVitala维塔拉.(*维塔拉VitalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vitala",
		TitleName: "维塔拉",
		TitleCode: "b_vitala",
	}
}
