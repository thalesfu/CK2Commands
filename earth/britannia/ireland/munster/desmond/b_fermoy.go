package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗莫伊FermoyBarony struct {
	feud.BaseBarony
}

var BFermoy弗莫伊 feud.Barony = &弗莫伊FermoyBarony{}

func init() {
    f := BFermoy弗莫伊.(*弗莫伊FermoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fermoy",
		TitleName: "弗莫伊",
		TitleCode: "b_fermoy",
	}
}
