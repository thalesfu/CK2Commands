package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希乌马HiiumaaBarony struct {
	feud.BaseBarony
}

var BHiiumaa希乌马 feud.Barony = &希乌马HiiumaaBarony{}

func init() {
	f := BHiiumaa希乌马.(*希乌马HiiumaaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hiiumaa",
		TitleName: "希乌马",
		TitleCode: "b_hiiumaa",
	}
}
