package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗克弗伊RoquefeuilBarony struct {
	feud.BaseBarony
}

var BRoquefeuil罗克弗伊 feud.Barony = &罗克弗伊RoquefeuilBarony{}

func init() {
    f := BRoquefeuil罗克弗伊.(*罗克弗伊RoquefeuilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roquefeuil",
		TitleName: "罗克弗伊",
		TitleCode: "b_roquefeuil",
	}
}
