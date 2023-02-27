package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉多堪波斯MarathokamposBarony struct {
	feud.BaseBarony
}

var BMarathokampos马拉多堪波斯 feud.Barony = &马拉多堪波斯MarathokamposBarony{}

func init() {
    f := BMarathokampos马拉多堪波斯.(*马拉多堪波斯MarathokamposBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marathokampos",
		TitleName: "马拉多堪波斯",
		TitleCode: "b_marathokampos",
	}
}
