package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞内泰什SzentesBarony struct {
	feud.BaseBarony
}

var BSzentes塞内泰什 feud.Barony = &塞内泰什SzentesBarony{}

func init() {
    f := BSzentes塞内泰什.(*塞内泰什SzentesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szentes",
		TitleName: "塞内泰什",
		TitleCode: "b_szentes",
	}
}
