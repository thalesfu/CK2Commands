package vinnytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩尔沃迈斯克PervomaiskBarony struct {
	feud.BaseBarony
}

var BPervomaisk佩尔沃迈斯克 feud.Barony = &佩尔沃迈斯克PervomaiskBarony{}

func init() {
    f := BPervomaisk佩尔沃迈斯克.(*佩尔沃迈斯克PervomaiskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pervomaisk",
		TitleName: "佩尔沃迈斯克",
		TitleCode: "b_pervomaisk",
	}
}
