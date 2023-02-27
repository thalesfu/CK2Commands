package caen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 栋夫龙DomfrontBarony struct {
	feud.BaseBarony
}

var BDomfront栋夫龙 feud.Barony = &栋夫龙DomfrontBarony{}

func init() {
    f := BDomfront栋夫龙.(*栋夫龙DomfrontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "domfront",
		TitleName: "栋夫龙",
		TitleCode: "b_domfront",
	}
}
