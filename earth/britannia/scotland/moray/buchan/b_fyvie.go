package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法维FyvieBarony struct {
	feud.BaseBarony
}

var BFyvie法维 feud.Barony = &法维FyvieBarony{}

func init() {
	f := BFyvie法维.(*法维FyvieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fyvie",
		TitleName: "法维",
		TitleCode: "b_fyvie",
	}
}
