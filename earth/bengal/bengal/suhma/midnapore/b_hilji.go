package midnapore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 醯梨恃HiljiBarony struct {
	feud.BaseBarony
}

var BHilji醯梨恃 feud.Barony = &醯梨恃HiljiBarony{}

func init() {
	f := BHilji醯梨恃.(*醯梨恃HiljiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hilji",
		TitleName: "醯梨恃",
		TitleCode: "b_hilji",
	}
}
