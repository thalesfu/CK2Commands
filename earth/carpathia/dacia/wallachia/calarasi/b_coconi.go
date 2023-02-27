package calarasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科科尼CoconiBarony struct {
	feud.BaseBarony
}

var BCoconi科科尼 feud.Barony = &科科尼CoconiBarony{}

func init() {
    f := BCoconi科科尼.(*科科尼CoconiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coconi",
		TitleName: "科科尼",
		TitleCode: "b_coconi",
	}
}
