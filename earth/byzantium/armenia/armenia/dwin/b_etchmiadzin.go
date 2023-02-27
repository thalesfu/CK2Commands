package dwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃奇米阿津EtchmiadzinBarony struct {
	feud.BaseBarony
}

var BEtchmiadzin埃奇米阿津 feud.Barony = &埃奇米阿津EtchmiadzinBarony{}

func init() {
    f := BEtchmiadzin埃奇米阿津.(*埃奇米阿津EtchmiadzinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "etchmiadzin",
		TitleName: "埃奇米阿津",
		TitleCode: "b_etchmiadzin",
	}
}
