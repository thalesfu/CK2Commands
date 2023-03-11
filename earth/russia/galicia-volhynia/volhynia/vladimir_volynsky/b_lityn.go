package vladimir_volynsky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利京LitynBarony struct {
	feud.BaseBarony
}

var BLityn利京 feud.Barony = &利京LitynBarony{}

func init() {
    f := BLityn利京.(*利京LitynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lityn",
		TitleName: "利京",
		TitleCode: "b_lityn",
	}
}
