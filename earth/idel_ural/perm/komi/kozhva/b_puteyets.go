package kozhva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普捷耶茨PuteyetsBarony struct {
	feud.BaseBarony
}

var BPuteyets普捷耶茨 feud.Barony = &普捷耶茨PuteyetsBarony{}

func init() {
    f := BPuteyets普捷耶茨.(*普捷耶茨PuteyetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puteyets",
		TitleName: "普捷耶茨",
		TitleCode: "b_puteyets",
	}
}
