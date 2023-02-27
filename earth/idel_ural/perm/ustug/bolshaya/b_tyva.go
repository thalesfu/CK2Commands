package bolshaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特瓦TyvaBarony struct {
	feud.BaseBarony
}

var BTyva特瓦 feud.Barony = &特瓦TyvaBarony{}

func init() {
    f := BTyva特瓦.(*特瓦TyvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyva",
		TitleName: "特瓦",
		TitleCode: "b_tyva",
	}
}
