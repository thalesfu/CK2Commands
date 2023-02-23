package bolshoy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛加派SingapayBarony struct {
	feud.BaseBarony
}

var BSingapay辛加派 feud.Barony = &辛加派SingapayBarony{}

func init() {
	f := BSingapay辛加派.(*辛加派SingapayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "singapay",
		TitleName: "辛加派",
		TitleCode: "b_singapay",
	}
}
