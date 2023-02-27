package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏穆波逊SymposionBarony struct {
	feud.BaseBarony
}

var BSymposion苏穆波逊 feud.Barony = &苏穆波逊SymposionBarony{}

func init() {
    f := BSymposion苏穆波逊.(*苏穆波逊SymposionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "symposion",
		TitleName: "苏穆波逊",
		TitleCode: "b_symposion",
	}
}
